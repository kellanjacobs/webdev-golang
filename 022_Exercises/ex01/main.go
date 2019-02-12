package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/cat/", jack)
	http.HandleFunc("/me/", about)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World")
}

func jack(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Jack the Cat")
}

func about(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Daniel's Page")
}