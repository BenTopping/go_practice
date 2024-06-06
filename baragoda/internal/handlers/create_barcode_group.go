package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BenTopping/go_practice/baragoda/api"
	"github.com/BenTopping/go_practice/baragoda/internal/tools"
	log "github.com/sirupsen/logrus"
)

func CreateBarcodeGroup(w http.ResponseWriter, r *http.Request) {
	var err error

	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	t := struct {
		Prefix *string `json:"prefix,omitempty"`
		Sequence *int `json:"sequence,string,omitempty"`
	}{}

	err = d.Decode(&t)

	if err != nil {
		api.RequestErrorHandler(w, fmt.Errorf("invalid JSON"))
		return
	}

	if t.Prefix == nil {
		api.RequestErrorHandler(w, fmt.Errorf("missing field 'prefix' from JSON object"))
		return
	}

	if t.Sequence == nil {
		api.RequestErrorHandler(w, fmt.Errorf("missing field 'sequence' from JSON object"))
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var barcodeGroup *tools.BarcodeGroup = (*database).CreateBarcodeGroup(*t.Prefix, *t.Sequence)
	if barcodeGroup == nil {
		api.InternalErrorHandler(w)
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