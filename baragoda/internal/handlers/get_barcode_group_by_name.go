package handlers

import (
	"net/http"
	"encoding/json"
	"fmt"

	"github.com/BenTopping/go_practice/baragoda/api"
	"github.com/BenTopping/go_practice/baragoda/internal/tools"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func GetBarcodeGroupByName(w http.ResponseWriter, r *http.Request) {
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

	var barcodeGroup *tools.BarcodeGroup = (*database).GetBarcodeGroup(barcodeGroupParam)
	if barcodeGroup == nil {
		api.NotFoundErrorHandler(w, fmt.Errorf("barcode group %s not found", barcodeGroupParam))
		return
	}

	var response = api.BarcodeGroupResponse{
		Code: http.StatusOK,
		Prefix: (*barcodeGroup).Prefix,
		Sequence: (*barcodeGroup).Sequence,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}