package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// Something returns some stuff for the website
func Something(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "AAAAA")
}

func main() {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("out"))
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	r.Handle("/", fs)
	r.HandleFunc("/api", Something)

	srv := &http.Server{
		Handler: r,
		Addr:    ":" + port,
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(srv.ListenAndServe())
}

func checkEnv() {
	var dbConnection string
	var storageConnection string

	env := os.Getenv("env")

	if strings.Compare(env, "prod") == 0 {
		fmt.Println("Running in Production")
	} else {
		fmt.Println("Running in Development")
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	dbConnection = os.Getenv("db_connection")
	storageConnection = os.Getenv("storage_connection")

	if dbConnection == "" {
		panic("Please make sure to define the db connection")
	}

	if storageConnection == "" {
		panic("Please make sure to define the storage connection")
	}
}
