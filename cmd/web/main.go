package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Priyansh-Kotak/booking-system/pkg/config"
	"github.com/Priyansh-Kotak/booking-system/pkg/handlers"
	"github.com/Priyansh-Kotak/booking-system/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8000"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.Secure = app.InProduction
	session.Cookie.SameSite = http.SameSiteLaxMode

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("failed to load the cache ")
	}

	app.UseCache = false
	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	errs := srv.ListenAndServe()
	log.Fatal(errs)

}
