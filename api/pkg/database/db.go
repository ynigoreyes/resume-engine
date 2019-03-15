package database

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/ynigoreyes/resume-engine/pkg/models"
)

// Create creates an instance of the database
func Create() *gorm.DB {
	log.Println("Creating new Google Cloud SQL connection...")

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

// Populate populates the database with the necessary users
func Populate(databaseName string, db *gorm.DB) {
	db.Exec(fmt.Sprintf("DROP table %v.users;", databaseName))
	db.AutoMigrate(&models.User{}, &models.Comment{})

	// Uncomment block to test Users table
	user1 := models.User{FirstName: "Chikorita", LastName: "Leaf Type", TagLine: "Just a Chikorita"}
	db.Create(&user1)
	user2 := models.User{FirstName: "Cyndaquil", LastName: "Fire Type", TagLine: "Just a Cyndaquil"}
	db.Create(&user2)
	user3 := models.User{FirstName: "Tododile", LastName: "Water Type", TagLine: "Just a Tododile"}
	db.Create(&user3)
	user4 := models.User{FirstName: "Eevee", LastName: "Normal Type", TagLine: "Just an Eevee"}
	db.Create(&user4)
}
