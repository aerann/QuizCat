package handlers

import (
	"net/http"

	chi "github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Check if it is a preflight request
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func Handler(r *chi.Mux) {
	// Global middleware
	r.Use(chimiddle.StripSlashes)
	r.Use(enableCors)

	r.Route("/generate_cards", func(r chi.Router) {
		r.Post("/", GenerateCards)
	},
	)

	r.Route("/", func(r chi.Router) {
		r.Get("/", HealthCheck)
	},
	)
}
