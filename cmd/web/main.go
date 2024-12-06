package main

import (
	"fmt"
	"github.com/Podmev/web-app-go/pkg/config"
	"github.com/Podmev/web-app-go/pkg/handlers"
	"github.com/Podmev/web-app-go/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Starting application on port %s", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)

}
