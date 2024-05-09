package handlers

import (
	"net/http"
	"encoding/json"
	"fmt"

	"github.com/BenTopping/go_practice/baragoda/api"
	"github.com/BenTopping/go_practice/baragoda/internal/tools"
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/schema"
)

func GetBarcode(w http.ResponseWriter, r *http.Request) {
	var params = api.BarcodeParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var barcodeSequence *tools.BarcodeSequence = (*database).GetBarcodeSequence(params.Barcode)
	if barcodeSequence == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var combinedBarcode string = fmt.Sprintf("%s-%d", (*barcodeSequence).Barcode, (*barcodeSequence).Sequence)

	var response = api.BarcodeResponse{
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