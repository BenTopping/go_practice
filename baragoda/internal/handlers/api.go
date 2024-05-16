package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)

	r.Route("/barcode_groups", func(r chi.Router) {

		r.Get("/", GetBarcodeGroups)

		r.Route("/{barcodeGroup}", func(r chi.Router) {
			r.Get("/", GetBarcode)
			r.Get("/new", CreateBarcode)
		})
	})
}