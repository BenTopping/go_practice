package tools

import (
	"database/sql"
	"log"
	"fmt"

	"github.com/go-sql-driver/mysql"
)


type Mysql struct {
	db *sql.DB
}

func (d *Mysql) SetupDatabase() (*Mysql, error) {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   "root",
		Passwd: "",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "baragoda_dev",
	}
	// Get a database handle.
	var err error
	var db *sql.DB
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	return &Mysql{ db }, err
}

func (m *Mysql) CreateBarcode(barcodeGroup string) (*BarcodeGroup, error) {
	var bg BarcodeGroup

	// Check the barcode group exists
	row := m.db.QueryRow("SELECT * FROM barcode_group WHERE prefix = ?", barcodeGroup)
	if err := row.Scan(&bg.Id, &bg.Prefix, &bg.Sequence, &bg.created_at); err != nil {
		if err == sql.ErrNoRows {
			// If no barcode group found, return an error
			return nil, fmt.Errorf("barcode group %s not found", barcodeGroup)
		}
		return &bg, fmt.Errorf("createBarcode %s: %v", barcodeGroup, err)
	}

	// Increase the sequence number
	bg.Sequence++
	// Update the sequence number in the database
	_, err := m.db.Exec("UPDATE barcode_group SET sequence = ? WHERE id = ?", bg.Sequence, bg.Id)
	if err != nil {
		// If update failed, return an error
		return nil, fmt.Errorf("unable to create new barcode")
	}

	return &bg, nil
}

func (d *Mysql) CreateBarcodeGroup(prefix string, sequence int) (*BarcodeGroup) {
	return nil
}

func (d *Mysql) GetBarcodeGroup(barcode string) (*BarcodeGroup) {
	return nil
}

func (d *Mysql) GetBarcodeGroups() (*[]BarcodeGroup) {
	return nil
}
