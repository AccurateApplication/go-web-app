package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}
