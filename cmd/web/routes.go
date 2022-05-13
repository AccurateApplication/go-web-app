package main

import (
	"net/http"

	"github.com/AccurateApplication/go-web-app/pkg/config"
	"github.com/AccurateApplication/go-web-app/pkg/handlers"
	"github.com/gorilla/mux"
)

func routes(conf *config.AppConf) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Repo.HelloHandler)
	r.Use(LoadSession)
	r.HandleFunc("/about", handlers.Repo.AboutHandler)
	r.Use(noSurf)
	return r
}
