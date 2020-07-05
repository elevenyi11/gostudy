package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"goji.io/pat"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func ping() {

	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error:", err)
	}

	fmt.Println("Connected to MongoDB!")
}

func collections() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("error:", err)
	}

	collection := client.Database("godb").Collection("col")

	log.Printf("data:%v", &collection)
}

type Person struct {
	Name  string
	Phone string
}

func mgov2() {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("godb").C("people")
	err = c.Insert(&Person{"superWang", "13478888311"}, &Person{"David", "15500001111"})
	if err != nil {
		log.Fatal(err)
	}
	result := Person{}
	err = c.Find(bson.M{"name": "superWang"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("name:", result.Name)
	fmt.Println("phone:", result.Phone)
}

type Book struct {
	ISBN    string   `json:"isbn"`
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
	Price   string   `json:"price"`
}

func ErrorWithJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "{message,%q}", message)
}
func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}

func EnsureIndex(s *mgo.Session) {
	session := s.Copy()
	defer session.Close()

	c := session.DB("store").C("books")

	index := mgo.Index{
		Key:        []string{"isbn"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}

func AllBooks(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		c := session.DB("store").C("books")

		var books []Book
		err := c.Find(bson.M{}).All(&books)
		if err != nil {
			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed get all books: ", err)
			return
		}

		respBody, err := json.MarshalIndent(books, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		ResponseWithJSON(w, respBody, http.StatusOK)
	}
}

func AddBook(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		var book Book
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&book)
		if err != nil {
			ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
			return
		}

		c := session.DB("store").C("books")

		err = c.Insert(book)
		if err != nil {
			if mgo.IsDup(err) {
				ErrorWithJSON(w, "Book with this ISBN already exists", http.StatusBadRequest)
				return
			}

			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed insert book: ", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Location", r.URL.Path+"/"+book.ISBN)
		w.WriteHeader(http.StatusCreated)
	}
}

func BookByISBN(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		isbn := pat.Param(r, "isbn")

		c := session.DB("store").C("books")

		var book Book
		err := c.Find(bson.M{"isbn": isbn}).One(&book)
		if err != nil {
			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed find book: ", err)
			return
		}

		if book.ISBN == "" {
			ErrorWithJSON(w, "Book not found", http.StatusNotFound)
			return
		}

		respBody, err := json.MarshalIndent(book, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		ResponseWithJSON(w, respBody, http.StatusOK)
	}
}

func UpdateBook(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		isbn := pat.Param(r, "isbn")

		var book Book
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&book)
		if err != nil {
			ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
			return
		}

		c := session.DB("store").C("books")

		err = c.Update(bson.M{"isbn": isbn}, &book)
		if err != nil {
			switch err {
			default:
				ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
				log.Println("Failed update book: ", err)
				return
			case mgo.ErrNotFound:
				ErrorWithJSON(w, "Book not found", http.StatusNotFound)
				return
			}
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func DeleteBook(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		isbn := pat.Param(r, "isbn")

		c := session.DB("store").C("books")

		err := c.Remove(bson.M{"isbn": isbn})
		if err != nil {
			switch err {
			default:
				ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
				log.Println("Failed delete book: ", err)
				return
			case mgo.ErrNotFound:
				ErrorWithJSON(w, "Book not found", http.StatusNotFound)
				return
			}
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
