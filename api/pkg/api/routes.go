package api

import (
	"encoding/json"
	"log"
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

// GetComments exposes the endpoint necessary for retriving resume comments
func (ro *Routes) GetComments(w http.ResponseWriter, r *http.Request) {
	// Extract route variables
	params := mux.Vars(r)

	var comments []struct {
		CommentBody        string `json:"comment_body"`
		CommenterID        uint   `json:"-"`
		CommenterFirstName string `json:"commenter_first_name"`
		CommenterLastName  string `json:"commenter_last_name"`
	}

	// Grabs all of the comments associated with the resume along with the user that commented
	ro.db.Raw("SELECT first_name, last_name, comment_body, commenter_id FROM users INNER JOIN comments ON comments.resume_id=users.id WHERE comments.resume_id=?", params["id"]).Scan(&comments)

	for index, comment := range comments {
		var user models.User
		ro.db.First(&user, comment.CommenterID)
		comments[index].CommenterFirstName = user.FirstName
		comments[index].CommenterLastName = user.LastName
	}

	// Return the result to the client
	w.Header().Set("Content-type", "application/json")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&comments)

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

	// Do insertion and report to client
	err = ro.db.Create(&comment).Error
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&comment)
	}

}
