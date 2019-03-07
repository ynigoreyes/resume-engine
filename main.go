package main

import (
	// "fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/ynigoreyes/resume-engine/pkg/api"
	"github.com/ynigoreyes/resume-engine/pkg/database"
	"github.com/ynigoreyes/resume-engine/pkg/models"
)

func main() {
	log.Printf("Initializing resume-engine...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	// Initialize database
	db := database.MustInit()
	defer db.Close()
	db.AutoMigrate(&models.User{}, &models.Comment{})
	/*
		// Uncomment block to test Users table
		user := models.User{UserID: "1", FirstName: "John", LastName: "Doe"}
		db.Create(&user)
		user = models.User{}
		fmt.Printf("%+v\n", db.First(&user, "user_id = ?", "1").Value)
	*/

	// Initialize API routes
	routes := api.CreateRoutes(db)

	// Define routes and API endpoints
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("out"))

	r.Handle("/", fs)
	r.HandleFunc("/api/comment/{user_id}", routes.CommentHandler).
		Methods("GET", "POST")

	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(srv.ListenAndServe())
}
