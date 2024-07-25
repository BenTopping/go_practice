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
		DBName: "goflag_dev",
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

func (m *Mysql) GetFlagGroups() (*[]FlagGroup, error) {
	var flagGroups []FlagGroup
	rows, err := m.db.Query("SELECT * FROM flag_group")

	if err != nil {
		return nil, fmt.Errorf("getFlagGroups: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var fg FlagGroup
		if err := rows.Scan(&fg.Id, &fg.Name, &fg.created_at); err != nil {
			return nil, fmt.Errorf("getFlagGroups: %v", err)
		}
		flagGroups = append(flagGroups, fg)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getFlagGroups: %v", err)
	}

	return &flagGroups, nil
}
