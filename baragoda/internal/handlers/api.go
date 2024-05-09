package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)

	r.Route("/barcodes_group", func(r chi.Router) {

		r.Get("/barcode", GetBarcode)
	})
}