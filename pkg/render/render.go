package render

import (
	"html/template"
	"log"
	"net/http"
)

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./templates/" + tmpl)
	if err != nil {
		log.Println("Error parsing files in templates. err: ", err)
	}
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("Error parsing template: ", err)
		return
	}
}
