package main

import (
	"net/http"

	"github.com/AccurateApplication/go-web-app/pkg/config"
	"github.com/AccurateApplication/go-web-app/pkg/handlers"
	"github.com/gorilla/mux"
)

func routes(conf *config.AppConf) http.Handler {
	r := mux.NewRouter()
	r.Use(LoadSession)
	r.Use(NoSurf)
	r.HandleFunc("/about", handlers.Repo.AboutHandler)
	r.HandleFunc("/", handlers.Repo.HelloHandler)
	fileServer := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/static").Handler(http.StripPrefix("/static", fileServer))
	return r
}
