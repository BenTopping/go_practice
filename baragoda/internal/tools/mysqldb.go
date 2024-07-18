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

func (m *Mysql) CreateBarcode(prefix string, count int) (*BarcodeGroup, error) {
	var bg BarcodeGroup

	// Check the barcode group exists
	row := m.db.QueryRow("SELECT * FROM barcode_group WHERE prefix = ?", prefix)
	if err := row.Scan(&bg.Id, &bg.Prefix, &bg.Sequence, &bg.created_at); err != nil {
		if err == sql.ErrNoRows {
			// If no barcode group found, return an error
			return nil, fmt.Errorf("barcode group %s not found", prefix)
		}
		return &bg, fmt.Errorf("createBarcode %s: %v", prefix, err)
	}

	// Increase the sequence number
	bg.Sequence += int64(count)
	// Update the sequence number in the database
	_, err := m.db.Exec("UPDATE barcode_group SET sequence = ? WHERE id = ?", bg.Sequence, bg.Id)
	if err != nil {
		// If update failed, return an error
		return nil, fmt.Errorf("unable to create new barcode")
	}

	return &bg, nil
}

func (m *Mysql) CreateBarcodeGroup(prefix string, sequence int) (*BarcodeGroup, error) {
	var bg BarcodeGroup

	_, err := m.db.Exec("INSERT INTO barcode_group (prefix, sequence) VALUES (?, ?)", prefix, sequence)
    if err != nil {
        return nil, fmt.Errorf("failed to create barcode group: %v", err)
    }

	row := m.db.QueryRow("SELECT * FROM barcode_group WHERE prefix = ?", prefix)
    if err := row.Scan(&bg.Id, &bg.Prefix, &bg.Sequence, &bg.created_at); err != nil {
		if err == sql.ErrNoRows {
			// If no barcode group found, return an error
			return nil, fmt.Errorf("barcode group %s not found", prefix)
		}
		return &bg, fmt.Errorf("createBarcodeGroup %s: %v", prefix, err)
	}

    return &bg, nil
}

func (m *Mysql) GetBarcodeGroup(prefix string) (*BarcodeGroup, error) {
	var bg BarcodeGroup

	// Check the barcode group exists
	row := m.db.QueryRow("SELECT * FROM barcode_group WHERE prefix = ?", prefix)
	if err := row.Scan(&bg.Id, &bg.Prefix, &bg.Sequence, &bg.created_at); err != nil {
		if err == sql.ErrNoRows {
			// If no barcode group found, return an error
			return nil, nil
		}
		return &bg, fmt.Errorf("getBarcodeGroup %s: %v", prefix, err)
	}

	return &bg, nil
}

func (m *Mysql) GetBarcodeGroups() (*[]BarcodeGroup, error) {
	var barcodeGroups []BarcodeGroup
	rows, err := m.db.Query("SELECT * FROM barcode_group")

	if err != nil {
		return nil, fmt.Errorf("getBarcodeGroups: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var bg BarcodeGroup
		if err := rows.Scan(&bg.Id, &bg.Prefix, &bg.Sequence, &bg.created_at); err != nil {
			return nil, fmt.Errorf("getBarcodeGroups: %v", err)
		}
		barcodeGroups = append(barcodeGroups, bg)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getBarcodeGroups: %v", err)
	}

	return &barcodeGroups, nil
}
