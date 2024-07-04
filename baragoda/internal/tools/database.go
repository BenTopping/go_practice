package tools

import (
	log "github.com/sirupsen/logrus"
)

type BarcodeGroup struct {
	Id int64
	Prefix string
	Sequence int64
	created_at []uint8
}

type DatabaseInterface interface {
	GetBarcodeGroups() (*[]BarcodeGroup, error)
	GetBarcodeGroup(barcode string) (*BarcodeGroup, error)
	CreateBarcode(barcode string) (*BarcodeGroup, error)
	CreateBarcodeGroup(prefix string, sequence int) (*BarcodeGroup, error)
	SetupDatabase() (*Mysql, error)
}

func NewDatabase() (*DatabaseInterface, error) {
	
	var database DatabaseInterface = &Mysql{}

	database, err := database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}