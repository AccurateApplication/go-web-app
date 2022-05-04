package main

import (
	"log"
	"net/http"

	"github.com/AccurateApplication/go-web-app/pkg/handlers"
)

const portNum = ":8080"

func main() {
	http.HandleFunc("/about", handlers.AboutHandler)
	http.HandleFunc("/", handlers.HelloHandler)
	log.Printf("Serving on %s", portNum)
	http.ListenAndServe(portNum, nil)

}
