package api

import (
	"net/http"
	"encoding/json"

	"github.com/BenTopping/go_practice/baragoda/internal/tools"
)

type BarcodeParams struct {
	Barcode string
}

type LastBarcodeResponse struct {
	Code int
	Barcode string
}

type CreateBarcodeResponse struct {
	Code int
	Barcodes []string
}

type BarcodeGroupResponse struct {
	Code int
	Prefix string
	Sequence int64
}

type BarcodeGroupsResponse struct {
	Code int
	BarcodeGroups []tools.BarcodeGroup
}

type Error struct {
	Code int
	Message string
}

func writeError(w http.ResponseWriter, message string, code int,) {
	resp := Error{
		Code: code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	NotFoundErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusNotFound)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An Unexpected Error Occured.", http.StatusInternalServerError)
	}
)