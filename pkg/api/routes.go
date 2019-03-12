package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/ynigoreyes/resume-engine/pkg/models"
)

// Routes encapsulates the dependencies necessary for defining API routes
type Routes struct {
	db *gorm.DB
}

// CreateRoutes initializes a new instance of the API route handlers
func CreateRoutes(db *gorm.DB) *Routes {
	return &Routes{db: db}
}

// GetComment exposes the endpoint necessary for retriving resume comments
func (ro *Routes) GetComment(w http.ResponseWriter, r *http.Request) {
	// Extract route variables
	params := mux.Vars(r)

	// Declare a comment to be referenced for storing query results
	var comment models.Comment

	// Get first comment entry from database that matches the requested ID
	err := ro.db.Where("id = ?", params["id"]).First(&comment).Error

	// Return the result to the client
	w.Header().Set("Content-type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&comment)
	}

}

// GetUsers exposes the endpoint necessary for getting users
func (ro *Routes) GetUsers(w http.ResponseWriter, r *http.Request) {
	// Declare a user to be referenced for storing query results
	var users []models.User

	// Get the users from the database
	err := ro.db.Find(&users).Error

	// Return the result to the client
	w.Header().Set("Content-type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&users)
	}

}

// AddComment exposes the endpoint necessary for adding comments
func (ro *Routes) AddComment(w http.ResponseWriter, r *http.Request) {
	// Declare a user to be referenced for storing query results
	var comment models.Comment

	// Decode request body into comment reference
	err := json.NewDecoder(r.Body).Decode(&comment)
	// defer r.Body.Close() // necessary?

	// Check if decode was successful
	w.Header().Set("Content-type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Check that request body is valid
	if validationErrors := comment.Validate(); len(validationErrors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": validationErrors})
		return
	}

	// Do insertion and report to client
	err = ro.db.Create(&comment).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&comment)
	}

}
