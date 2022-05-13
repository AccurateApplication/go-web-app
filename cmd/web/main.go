package main

import (
	"log"
	"net/http"
	"time"

	"github.com/AccurateApplication/go-web-app/pkg/config"
	"github.com/AccurateApplication/go-web-app/pkg/handlers"
	"github.com/AccurateApplication/go-web-app/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNum = ":8080"

var conf config.AppConf
var session *scs.SessionManager

func main() {

	conf.SecureSite = false

	// Initialize a new session manager and configure the session lifetime.
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = conf.SecureSite

	conf.Session = session

	tc, err := render.CreateTmplCache()
	if err != nil {
		log.Fatalln("Error creating template cache: ", err)
	}
	conf.TemplateCache = tc
	render.GenerateTemplates(&conf)

	repo := handlers.NewRepo(&conf)
	handlers.NewHandlers(repo)

	s := &http.Server{
		Handler: routes(&conf),
		Addr:    portNum,
	}
	log.Printf("Serving on %s", portNum)
	log.Fatal(s.ListenAndServe())
}
