package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresConn(databaseUrl string) (*gorm.DB, error) {
	// Connect to the database
	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to the postgres database")
	return db, nil
}