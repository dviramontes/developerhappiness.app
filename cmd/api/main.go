package main

import (
	"fmt"
	"github.com/dviramontes/developerhappiness.app/pkg/api"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"net/http"
	"os"
	"time"
)

func main() {
	var port string

	port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	API := api.New()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// basic CORS
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	// middleware setup
	r.Use(
		cors.Handler,
		render.SetContentType(render.ContentTypeJSON), // set content-type headers as application/json
		middleware.Logger,                         // log api request calls
		middleware.StripSlashes,                   // match paths with a trailing slash, strip it, and continue routing through the mux
		middleware.Recoverer,                      // recover from panics without crashing server
		middleware.Timeout(3000*time.Millisecond), // Stop processing after 3 seconds
	)

	// obligatory health-check endpoint
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/webhook", func(r chi.Router) {
			r.Post("/slack", API.SlackHandler)
		})
	})

	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
