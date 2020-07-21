package net

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"
)
func SayHello(w http.ResponseWriter, req *http.Request){
	w.Write([]byte("hello"))
}

func TestHttp(t *testing.T){
	http.HandleFunc("/hello", SayHello)
	http.ListenAndServe(":8001",nil)
}
