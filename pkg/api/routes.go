package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/ynigoreyes/resume-engine/pkg/models"
	"net/http"
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
	w.Header().Set("Content-type", "applciation/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&comment)
	}

}

// GetUser exposes the endpoint necessary for getting users
func (ro *Routes) GetUser(w http.ResponseWriter, r *http.Request) {
	// Extract route variables
	params := mux.Vars(r)

	// Declare a user to be referenced for storing query results
	var user models.User

	// Get first user entry from database that matches the requested ID
	err := ro.db.Where("id = ?", params["id"]).First(&user).Error

	// Return the result to the client
	w.Header().Set("Content-type", "applciation/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&user)
	}

}

// AddComment exposes the endpoint necessary for adding comments
func (ro *Routes) AddComment(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Status       int
		ResponseBody string
	}{
		Status:       http.StatusOK,
		ResponseBody: "Successfully added comment.",
	}
	json.NewEncoder(w).Encode(data)
}
