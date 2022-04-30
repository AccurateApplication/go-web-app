package main

import (
	"fmt"
	"html/template"
	"log"
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

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("Error parsing template: ", err)
		return
	}
}
