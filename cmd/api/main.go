package main

import (
	"fmt"
	"github.com/dviramontes/developerhappiness.app/internal/config"
	"github.com/dviramontes/developerhappiness.app/pkg/api"
	"github.com/dviramontes/developerhappiness.app/pkg/db"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"

	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	var port string

	port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	conf := config.Read("config.yaml", nil)

	version := conf.GetString("version")
	connStr := conf.GetString("connStr")

	log.Printf("version :: %s", version)

	database, err := db.Connect(connStr)
	if err != nil {
		log.Fatalf("failed to connect to database, err: %v", err)
	}
	defer database.Conn.Close()

	API := api.New(conf, database)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// basic CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	// middleware setup
	r.Use(
		corsHandler.Handler,
		render.SetContentType(render.ContentTypeJSON), // set content-type headers as application/json
		middleware.Logger,                             // log api request calls
		middleware.StripSlashes,                       // match paths with a trailing slash, strip it, and continue routing through the mux
		middleware.Recoverer,                          // recover from panics without crashing server
		middleware.Timeout(3000*time.Millisecond),     // Stop processing after 3 seconds
	)

	// obligatory health-check endpoint
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/webhook", func(r chi.Router) {
			r.Post("/slack", API.SlackHandler)
		})
		r.Get("/users", API.GetUsers)
	})

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "client/build"))
	FileServer(r, "/", filesDir)

	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileSever does not allow any URL parameters")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}

	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
