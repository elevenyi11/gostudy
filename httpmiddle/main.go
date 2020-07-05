package main

import (
	"net/http"
	"net/http/httptest"
)

// type SingleHost struct {
// 	handler     http.Handler
// 	allowedHost string
// }

// func (this *SingleHost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	println(r.Host)
// 	if r.Host == this.allowedHost {
// 		this.handler.ServeHTTP(w, r)
// 	} else {
// 		w.WriteHeader(403)
// 	}

// }

// func myHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("hello word"))
// }
// func main() {

// 	single := &SingleHost{
// 		handler:     http.HandlerFunc(myHandler),
// 		allowedHost: "localhost:8080",
// 	}

// 	http.ListenAndServe(":8080", single)
// }

// func SingleHost(handler http.Handler, allowedHost string) http.Handler {
// 	fn := func(w http.ResponseWriter, r *http.Request) {
// 		if r.Host == allowedHost {
// 			handler.ServeHTTP(w, r)
// 		} else {
// 			w.WriteHeader(403)
// 		}
// 	}
// 	return http.HandlerFunc(fn)
// }

// func myHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("hello word"))
// }

// func main() {
// 	single := SingleHost(http.HandlerFunc(myHandler), "localhost:8080")
// 	http.ListenAndServe(":8080", single)
// }

// type AppendMiddleware struct {
// 	handler http.Handler
// }

// func (this *AppendMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	this.handler.ServeHTTP(w, r)
// 	w.Write([]byte(" hello world, this is middleware\n"))
// }

// func myHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte(" hello my handler\n"))
// }

// func main() {

// 	mid := &AppendMiddleware{http.HandlerFunc(myHandler)}
// 	http.ListenAndServe(":8080", mid)
// }

type ModifierMiddleware struct {
	handler http.Handler
}

func (this *ModifierMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rec := httptest.NewRecorder()
	this.handler.ServeHTTP(rec, r)

	for k, v := range rec.Header() {
		w.Header()[k] = v
	}
	w.Header().Set("go-web-foundation", "vip")
	w.WriteHeader(418)
	w.Write(rec.Body.Bytes())
	w.Write([]byte("hello world\n"))
	w.Write([]byte("hey this is middleware\n"))
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello this is my handler\n"))
}

func main() {

	mid := &ModifierMiddleware{http.HandlerFunc(myHandler)}
	http.ListenAndServe(":8080", mid)
}
