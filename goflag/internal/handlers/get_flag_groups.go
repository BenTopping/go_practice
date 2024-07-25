package handlers

import (
	"net/http"
	"encoding/json"

	"github.com/BenTopping/go_practice/goflag/api"
	"github.com/BenTopping/go_practice/goflag/internal/tools"
	log "github.com/sirupsen/logrus"
)

func GetFlagGroups(w http.ResponseWriter, r *http.Request) {
	var err error

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	flagGroups, err := (*database).GetFlagGroups()

	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var response = api.FlagGroupsResponse{
		Code: http.StatusOK,
		FlagGroups: *flagGroups,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
