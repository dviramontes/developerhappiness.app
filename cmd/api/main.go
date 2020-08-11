package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"net/http"
	"time"
)

type SlackURLVerifyPayload struct {
	Token     string `json:"token"`
	Challenge string `json:"challenge"`
	Type      string `json:"type"`
}

type SlackURLVerifyResponse struct {
	Challenge string `json:"challenge"`
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Basic CORS
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
		r.Post("/webhook", func(w http.ResponseWriter, r *http.Request) {
			var sURLVerify SlackURLVerifyPayload
			if err := json.NewDecoder(r.Body).Decode(&sURLVerify); err != nil {
				http.Error(w, "error decoding json payload from slack verify webhook", http.StatusInternalServerError)
				return
			}
			// respond with JSON Challenge
			res := SlackURLVerifyResponse{sURLVerify.Challenge}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(res)
		})
	})

	http.ListenAndServe(":3000", r)
}
