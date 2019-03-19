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

	log.Printf("Connecting to the %v database", databaseName)

	db, err := gorm.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	return db
}

// Populate populates the database with the necessary users
func Populate(databaseName string, db *gorm.DB) {
	db.Exec(fmt.Sprintf("DROP table %v.users;", databaseName))
	db.Exec(fmt.Sprintf("DROP table %v.comments;", databaseName))
	db.AutoMigrate(&models.User{}, &models.Comment{})

	// Uncomment block to test Users table
	user1 := models.User{FirstName: "Chikorita", LastName: "Grass Type", TagLine: "Just a Chikorita"}
	db.Create(&user1)
	user2 := models.User{FirstName: "Cyndaquil", LastName: "Fire Type", TagLine: "Just a Cyndaquil"}
	db.Create(&user2)
	user3 := models.User{FirstName: "Tododile", LastName: "Water Type", TagLine: "Just a Tododile"}
	db.Create(&user3)
	user4 := models.User{FirstName: "Eevee", LastName: "Normal Type", TagLine: "Just an Eevee"}
	db.Create(&user4)
	googler := models.User{FirstName: "Google", LastName: "Student", TagLine: "A Texas Tech Student"}
	db.Create(&googler)

	comment1 := models.Comment{ResumeID: uint(user1.ID), CommenterID: uint(user2.ID), CommentBody: "I think it's okay"}
	db.Create(&comment1)
	comment2 := models.Comment{ResumeID: uint(user1.ID), CommenterID: uint(user3.ID), CommentBody: "This is great! Keep up the good work!"}
	db.Create(&comment2)
	comment3 := models.Comment{ResumeID: uint(user1.ID), CommenterID: uint(user4.ID), CommentBody: "This is really bad"}
	db.Create(&comment3)

	comment4 := models.Comment{ResumeID: uint(user2.ID), CommenterID: uint(user1.ID), CommentBody: "I love it so much!"}
	db.Create(&comment4)

	comment5 := models.Comment{ResumeID: uint(user3.ID), CommenterID: uint(user1.ID), CommentBody: "I'm not really sure how I feel about all of this..."}
	db.Create(&comment5)
	comment6 := models.Comment{ResumeID: uint(user3.ID), CommenterID: uint(user2.ID), CommentBody: "This could be better"}
	db.Create(&comment6)
	comment7 := models.Comment{ResumeID: uint(user3.ID), CommenterID: uint(user4.ID), CommentBody: "I expected a lot more from you"}
	db.Create(&comment7)

	comment8 := models.Comment{ResumeID: uint(user4.ID), CommenterID: uint(user1.ID), CommentBody: "This sounds very professional"}
	db.Create(&comment8)
	comment9 := models.Comment{ResumeID: uint(user4.ID), CommenterID: uint(user3.ID), CommentBody: "I am not impressed. I'm sorry"}
	db.Create(&comment9)
}
