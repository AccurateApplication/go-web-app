package main

import (
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}
