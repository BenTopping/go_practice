package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)

	r.Route("/barcode_groups", func(r chi.Router) {

		r.Get("/", GetBarcodeGroups)
		r.Post("/new", CreateBarcodeGroup)

		r.Route("/{barcodeGroup}", func(r chi.Router) {
			r.Get("/", GetBarcodeGroupByName)
		})
	})

	r.Route("/barcodes", func(r chi.Router) {
		r.Route("/{prefix}", func(r chi.Router) {
			r.Post("/new", CreateBarcode)
			r.Get("/last", GetLastBarcode)
		})
	})
}