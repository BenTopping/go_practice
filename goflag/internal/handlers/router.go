package handlers

import (
	"github.com/go-chi/chi"
	"github.com/BenTopping/go_practice/goflag/internal/handlers/api"
	"github.com/BenTopping/go_practice/goflag/internal/handlers/ui"

	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)

	r.Route("/", func(r chi.Router) {
		r.Get("/", ui.GetIndexPage)
	})

	r.Route("/api", func(r chi.Router) {
		r.Get("/", api.GetFlagGroups)
	})
}
