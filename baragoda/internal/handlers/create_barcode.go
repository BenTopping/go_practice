package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BenTopping/go_practice/baragoda/api"
	"github.com/BenTopping/go_practice/baragoda/internal/tools"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func CreateBarcode(w http.ResponseWriter, r *http.Request) {
	var err error

	prefixParam := chi.URLParam(r, "prefix")

	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	t := struct {
		Count *int `json:"count,string,omitempty"`
	}{}

	err = d.Decode(&t)

	if err != nil {
		api.RequestErrorHandler(w, fmt.Errorf("invalid JSON"))
		return
	}

	if prefixParam == "" {
		api.RequestErrorHandler(w, fmt.Errorf("a barcode group prefix is required"))
		return
	}

	// If count is nil set it to 1
	if t.Count == nil {
		api.RequestErrorHandler(w, fmt.Errorf("a count param is required"))
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	barcodeGroup, err := (*database).CreateBarcode(prefixParam, *t.Count)
	if err != nil {
		api.NotFoundErrorHandler(w, err)
		return
	}

	barcodes := []string{}
	for i := 0; i < *t.Count; i++ {
		currentSequence := barcodeGroup.Sequence - int64(*t.Count) + int64(i) + 1
		combinedBarcode := fmt.Sprintf("%s-%d", barcodeGroup.Prefix, currentSequence)
		barcodes = append(barcodes, combinedBarcode)
	}

	var response = api.CreateBarcodeResponse{
		Code:     http.StatusOK,
		Barcodes: barcodes,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
