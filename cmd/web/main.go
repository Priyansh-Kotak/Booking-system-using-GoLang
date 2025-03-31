package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Priyansh-Kotak/udemy-course-project/pkg/config"
	"github.com/Priyansh-Kotak/udemy-course-project/pkg/handlers"
	"github.com/Priyansh-Kotak/udemy-course-project/pkg/render"
)

const portNumber = ":8000"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("failed to load the cache ")
	}

	app.UseCache = false
	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("Starting port at ", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
