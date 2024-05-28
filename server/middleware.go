package server

import (
	"net/http"
	"notes-app/internal/config"

	"github.com/justinas/nosurf"
)

// NoSurf is the csrf protection middleware
func NoSurf(app *config.AppConfig) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		csrfHandler := nosurf.New(next)
		csrfHandler.SetBaseCookie(http.Cookie{
			HttpOnly: true,
			Path:     "/",
			Secure:   app.InProduction,
			SameSite: http.SameSiteLaxMode,
		})
		return csrfHandler
	}
}

func SessionLoad(app *config.AppConfig) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return app.Session.LoadAndSave(next)
	}
}
