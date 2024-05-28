package main

import (
	"fmt"
	"log"
	"net/http"
	"notes-app/internal/config"
	"notes-app/internal/handlers"
	"notes-app/internal/render"
	"notes-app/server"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig

// main is the main function
func main() {

	app.InProduction = false
	fmt.Println("creating cache here... ")
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	// this library is for maintainng sessions
	session := scs.New()
	// keeps session alive
	session.Lifetime = 24 * 60 * 60
	// persist cookie even on re-load
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Starting application on port %s\n", portNumber)

	srv := http.Server{
		Addr:    portNumber,
		Handler: server.Routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
