package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
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
