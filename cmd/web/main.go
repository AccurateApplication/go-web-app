package main

import (
	"log"
	"net/http"

	"github.com/AccurateApplication/go-web-app/pkg/config"
	"github.com/AccurateApplication/go-web-app/pkg/handlers"
	"github.com/AccurateApplication/go-web-app/pkg/render"
)

const portNum = ":8080"

func main() {
	var conf config.AppConf

	tc, err := render.CreateTmplCache()
	if err != nil {
		log.Fatalln("Error creating template cache: ", err)
	}
	conf.TemplateCache = tc
	render.GenerateTemplates(&conf)

	repo := handlers.NewRepo(&conf)
	handlers.NewHandlers(repo)

	http.HandleFunc("/about", handlers.Repo.AboutHandler)
	http.HandleFunc("/", handlers.Repo.HelloHandler)
	log.Printf("Serving on %s", portNum)
	http.ListenAndServe(portNum, nil)

}
