package main

import (
	"io"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!")
}
func helloHandler11(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!Hello, world!")
}
func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/world", helloHandler11)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
