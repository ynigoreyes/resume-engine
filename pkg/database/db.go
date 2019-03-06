package database

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

// MustInit initializes an instance of the database
func MustInit() *gorm.DB {
	log.Println("Initializing Cloud SQL connection...")

	// Extract database credentials from environment
	username := os.Getenv("SQL_USERNAME")
	password := os.Getenv("SQL_PASSWORD")
	connectionName := os.Getenv("SQL_CONNECTION_NAME")
	databaseName := os.Getenv("SQL_DATABASE")

	if username == "" {
		panic(errors.New("SQL_USERNAME var is empty"))
	}
	if password == "" {
		panic(errors.New("SQL_PASSWORD var is empty"))
	}
	if connectionName == "" {
		panic(errors.New("SQL_CONNECTION_NAME var is empty"))
	}
	if databaseName == "" {
		panic(errors.New("SQL_DATABASE var is empty"))
	}
	dataSourceName := fmt.Sprintf("%s:%s@cloudsql(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, connectionName, databaseName)

	db, err := gorm.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	return db
}
