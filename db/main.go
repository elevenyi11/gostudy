package main

import (
	"net/http"

	"goji.io"
	"goji.io/pat"
	"gopkg.in/mgo.v2"
)

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	EnsureIndex(session)

	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/books"), AllBooks(session))
	mux.HandleFunc(pat.Post("/books"), AddBook(session))
	mux.HandleFunc(pat.Get("/books/:isbn"), BookByISBN(session))
	mux.HandleFunc(pat.Put("/books/:isbn"), UpdateBook(session))
	mux.HandleFunc(pat.Delete("/books/:isbn"), DeleteBook(session))
	http.ListenAndServe("localhost:8080", mux)

}
