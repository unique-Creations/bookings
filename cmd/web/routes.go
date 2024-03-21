package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/unique-Creations/bookings/config"
	"github.com/unique-Creations/bookings/pkg/handlers"
	"net/http"
)

// routes defines the routes for the application
func routes(app *config.AppConfig) http.Handler {
	//mux := pat.New()
	//mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	//mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	//mux.Use(writeToConsole)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
