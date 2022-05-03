package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
	fmt.Fprintf(w, "Hello")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
	fmt.Fprintf(w, "this is about page")
}
