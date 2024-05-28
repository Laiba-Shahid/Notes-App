package server

import (
	"net/http"
	"notes-app/internal/config"
	"notes-app/internal/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes(config *config.AppConfig) http.Handler {

	//chi is a lightweight, idiomatic and composable router for building Go HTTP services.
	mux := chi.NewRouter()

	mux.Use(NoSurf(config))
	mux.Use(middleware.Recoverer)
	mux.Use(SessionLoad(config))
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux

}
