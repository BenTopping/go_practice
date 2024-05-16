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

	barcodeGroupParam := chi.URLParam(r, "barcodeGroup")

	if barcodeGroupParam == "" {
		api.RequestErrorHandler(w, fmt.Errorf("a barcode group is required"))
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var barcodeGroup *tools.BarcodeGroup = (*database).CreateBarcode(barcodeGroupParam)
	if barcodeGroup == nil {
		api.NotFoundErrorHandler(w, fmt.Errorf("barcode group %s not found", barcodeGroupParam))
		return
	}

	var combinedBarcode string = fmt.Sprintf("%s-%d", (*barcodeGroup).Barcode, (*barcodeGroup).Sequence)

	var response = api.CreateBarcodeResponse{
		Code: http.StatusOK,
		Barcode: combinedBarcode,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}