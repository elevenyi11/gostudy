// package main

// import (
// 	"io"
// 	"log"
// 	"net/http"
// )

// func main1() {
// 	http.HandleFunc("/", sayHello)
// 	err := http.ListenAndServe(":8080", nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func sayHello(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, "hello world, this is version 1.")
//}
