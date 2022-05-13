package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf adds CSRF protection as middleware
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   conf.SecureSite,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// Load and save session on every request
func LoadSession(next http.Handler) http.Handler {
	return session.LoadAndSave(next)

}
