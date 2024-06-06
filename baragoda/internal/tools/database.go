package tools

import (
	log "github.com/sirupsen/logrus"
)

type BarcodeGroup struct {
	Prefix string
	Sequence int
}

type DatabaseInterface interface {
	GetBarcodeGroups() *[]BarcodeGroup
	GetBarcodeGroup(barcode string) *BarcodeGroup
	CreateBarcode(barcode string) *BarcodeGroup
	CreateBarcodeGroup(prefix string, sequence int) *BarcodeGroup
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