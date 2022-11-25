package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/matheus1103/go-studies/pkg/config"
	"github.com/matheus1103/go-studies/pkg/handlers"
	"github.com/matheus1103/go-studies/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {

	//TODO: change this to true when in production
	app.InProduction = false

	// initialize session
	session = scs.New()
	session.Lifetime = 24 * 60 // 24 hours
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("starting application on port %s", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
