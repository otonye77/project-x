package main

import (
	"fmt"
	"log"
	"net/http"
	"main.com/mymodule/pkg/config"
	"main.com/mymodule/pkg/handlers"
	"main.com/mymodule/pkg/render"
)

const portNumber = ":8080"

func main(){
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tc
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
