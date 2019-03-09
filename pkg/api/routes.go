package api

import (
	"encoding/json"
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
	json.NewEncoder(w).Encode(&models.Comment{})

}

// GetUser exposes the endpoint necessary for getting users
func (ro *Routes) GetUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(&models.User{})

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
