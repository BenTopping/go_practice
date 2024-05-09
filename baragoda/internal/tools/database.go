package tools

import (
	log "github.com/sirupsen/logrus"
)

type BarcodeSequence struct {
	Barcode string
	Sequence int
}

type DatabaseInterface interface {
	GetBarcodeSequence(barcode string) *BarcodeSequence
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {
	
	var database DatabaseInterface = &mockDB{}

	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}