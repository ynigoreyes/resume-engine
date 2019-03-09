package main

import (
	// "fmt"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
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
		IdleTimeout:  60 * time.Second,
	}

	// Run server in a go routine so it doesn't block
	go func() {
		log.Printf("Listening on port: %v", port)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal
	<-c

	// Create a deadline to wait 1 minute for existing connections to terminate
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline
	srv.Shutdown(ctx)
	log.Println("Shutting down resume-engine...")
	os.Exit(0)

}
