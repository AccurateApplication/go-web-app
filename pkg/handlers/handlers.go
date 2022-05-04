package handlers

import (
	"net/http"

	"github.com/AccurateApplication/go-web-app/pkg/render"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
