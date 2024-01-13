package handlers

import (
	chi "github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	// Global middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/generate_cards", func(r chi.Router) {
			r.Get("/", GenerateCards)
		},
	)

	r.Route("/", func(r chi.Router) {
			r.Get("/", Heartbeat)
		},
	)
}