package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/unique-Creations/bookings/config"
	"github.com/unique-Creations/bookings/pkg/handlers"
	"github.com/unique-Creations/bookings/pkg/render"
	"log"
	"net/http"
	"time"
)

const PORTNUMBER = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	// Change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour              // session last for 24 hours
	session.Cookie.Persist = true                  // session persist after browser window is closed
	session.Cookie.SameSite = http.SameSiteLaxMode // session cookie is sent with every request
	session.Cookie.Secure = app.InProduction       // session cookie is not sent over https
	app.Session = session                          // set the session to the app config

	tc, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Println("error creating template cache:", err.Error())
	}
	app.TemplateCache = tc
	app.UseCache = false
	render.NewTemplates(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	fmt.Printf("Starting app on port %s", PORTNUMBER)
	srv := &http.Server{
		Addr:    PORTNUMBER,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
