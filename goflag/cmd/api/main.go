package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/BenTopping/go_practice/goflag/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	// Handle static assets
	r.Handle("/css/*", http.StripPrefix("/css/", http.FileServer(http.Dir("./templates/css"))))
	r.Handle("/js/*", http.StripPrefix("/js/", http.FileServer(http.Dir("./templates/js"))))

	fmt.Println("Starting server on port 8080")

	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Error(err)
	}
}
