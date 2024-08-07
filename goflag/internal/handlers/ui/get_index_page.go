package ui

import (
	"net/http"
	"html/template"

	"github.com/BenTopping/go_practice/goflag/api"
	"github.com/BenTopping/go_practice/goflag/internal/tools"
	log "github.com/sirupsen/logrus"
)

func GetIndexPage(w http.ResponseWriter, r *http.Request) {
	var err error

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	flagGroups, err := (*database).GetFlagGroups()

	for i, flagGroup := range *flagGroups {
		(*flagGroups)[i].Flags, err = (*database).GetFlags(flagGroup.Name)
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}
	}
		
	data := map[string]interface{}{
		"FlagGroups": flagGroups,
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Error(err)
		return
	}
	tmpl.Execute(w, data)
}
