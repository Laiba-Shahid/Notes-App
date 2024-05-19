package main

import (
	"fmt"
	"log"
	"net/http"
	"notes-app/internal/config"
	"notes-app/internal/handlers"
	"notes-app/internal/render"
)

const portNumber = ":8080"

// main is the main function
func main() {
	var app config.AppConfig

	fmt.Println("creating cache here... ")
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
