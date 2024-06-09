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

	// Create a file server that serves files from the "./static/" directory.
	fileserver := http.FileServer(http.Dir("./static/"))

	// Handle requests with the "/static/*" pattern.
	// Strip the "/static" prefix from the request URL and pass it to the file server.
	mux.Handle("/static/*", http.StripPrefix("/static", fileserver))

	//routes for the application
	mux.Get("/", handlers.Repo.Home)

	mux.Get("/signup", handlers.Repo.SignUp)

	mux.Get("/login", handlers.Repo.Login)
	mux.Post("/login", handlers.Repo.PostLogin)

	return mux

}
