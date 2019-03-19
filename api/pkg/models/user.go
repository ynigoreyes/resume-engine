package models

import (
	"github.com/jinzhu/gorm"
)

// User defines a user interacting with the application
type User struct {
	gorm.Model
	FirstName string `json:"first_name,omitempty" gorm:"not null;type:varchar(25)"`
	LastName  string `json:"last_name,omitempty" gorm:"not null;type:varchar(25)"`
	TagLine   string `json:"tag_line,omitempty" gorm:"not null;type:varchar(25)"`
}
