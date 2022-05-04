package handlers

import (
	"net/http"

	"github.com/AccurateApplication/go-web-app/pkg/config"
	"github.com/AccurateApplication/go-web-app/pkg/render"
)

// Repo the respository that handlers use
var Repo *Repository

type Repository struct {
	Conf *config.AppConf
}

// Create new repo
func NewRepo(c *config.AppConf) *Repository {
	return &Repository{
		Conf: c,
	}
}

// Set repo for handlers
func NewHandlers(r *Repository) {
	Repo = r

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
