package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

func noSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   conf.SecureSite,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}
