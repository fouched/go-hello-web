package main

import (
	"fmt"
	"github.com/fouched/go-course/pkg/config"
	"github.com/fouched/go-course/pkg/handlers"
	"github.com/fouched/go-course/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc

	render.NewTemplates(&app)

	// pass a function to / that will handle requests
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on %s", portNumber))
	// ignore errors
	_ = http.ListenAndServe(portNumber, nil)
}
