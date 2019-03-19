package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/ynigoreyes/resume-engine/pkg/api"
	"github.com/ynigoreyes/resume-engine/pkg/database"
)

func main() {
	log.Printf("Initializing resume-engine...")

	// For the sake of the Demo, don't ignore the .env file
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
	db := database.Create()
	defer db.Close()
	// Will populate the database for me
	database.Populate(os.Getenv("SQL_DATABASE"), db)

	// Initialize API routes
	routes := api.CreateRoutes(db)

	// Define routes and API endpoints
	r := mux.NewRouter()

	// /api routes allow the React frontent to communicate with the backend
	r.HandleFunc("/api/user/{id}", routes.GetUser).Methods("GET")
	r.HandleFunc("/api/comment/{id}", routes.GetComments).Methods("GET")
	r.HandleFunc("/api/user", routes.GetUsers).Methods("GET")
	r.HandleFunc("/api/comment", routes.AddComment).Methods("POST")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})

	// Create an http server using the mux router
	srv := &http.Server{
		Handler:      handlers.CORS(headersOk)(r),
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
