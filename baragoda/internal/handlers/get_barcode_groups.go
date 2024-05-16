package handlers

import (
	"net/http"
	"encoding/json"

	"github.com/BenTopping/go_practice/baragoda/api"
	"github.com/BenTopping/go_practice/baragoda/internal/tools"
	log "github.com/sirupsen/logrus"
)

func GetBarcodeGroups(w http.ResponseWriter, r *http.Request) {
	var err error

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var barcodeGroups *[]tools.BarcodeGroup = (*database).GetBarcodeGroups()

	var response = api.BarcodeGroupsResponse{
		Code: http.StatusOK,
		BarcodeGroups: *barcodeGroups,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}