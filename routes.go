package main

import (
	"net/http"

	"goji.io"
	"goji.io/pat"

	"github.com/pykmi/api-seed/datastore"
	"github.com/pykmi/api-seed/handlers"
)

// corsMiddle middleware handles the supported access-control headers
func corsMiddle(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS, DELETE")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func setupRoutes(store *datastore.Store) http.Handler {
	// API Routes
	router := goji.NewMux()

	router.Use(corsMiddle)
	router.Use(datastore.Middleware(store))
	//router.Use(auth.TokenMiddleware)

	router.HandleFunc(pat.Get("/"), handlers.Default)

	return router
}
