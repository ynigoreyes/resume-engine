package models

import (
	"github.com/jinzhu/gorm"
)

// User defines a user interacting with the application
type User struct {
	gorm.Model
	UserID    string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	TagLine   string `json:"tag_line"`
	ResumeID  string `json:"resume_id"`
}
