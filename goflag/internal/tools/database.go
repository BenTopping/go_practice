package tools

import (
	log "github.com/sirupsen/logrus"
)

type Flag struct {
	Id int64
	Name string
	Description string
	Enabled bool
	created_at []uint8
	flag_group_id int64
}

type FlagGroup struct {
	Id int64
	Name string
	created_at []uint8
	Flags *[]Flag
}

// DatabaseInterface is an interface that defines the methods that a database must implement
type DatabaseInterface interface {
	GetFlagGroups() (*[]FlagGroup, error)
	GetFlags(flag_group_name string) (*[]Flag, error)
	SetupDatabase() (*Mysql, error)
}

// NewDatabase is a function that returns a database connection given an interface
func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &Mysql{}

	database, err := database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}
