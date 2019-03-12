package models

import (
	"github.com/jinzhu/gorm"
	"net/url"
)

// User defines a user interacting with the application
type User struct {
	gorm.Model
	FirstName string `json:"first_name,omitempty" gorm:"type:varchar(25)"`
	LastName  string `json:"last_name,omitempty" gorm:"type:varchar(25)"`
	TagLine   string `json:"tag_line,omitempty" gorm:"type:varchar(25)"`
}

// Validate validates a user creation request
func (u *User) Validate() url.Values {
	// Track request body errors for response
	errs := url.Values{}

	// Check if the first name empty
	if u.FirstName == "" {
		errs.Add("first_name", "first_name field is required")
	}

	// Check if the last name empty
	if u.FirstName == "" {
		errs.Add("last_name", "last_name field is required")
	}

	// Check if the tag line is empty
	if u.TagLine == "" {
		errs.Add("tag_line", "tag_line field is required")
	}

	// Check if the tag line is too long
	if len(u.TagLine) > 25 {
		errs.Add("tag_line", "tag_line cannot be greater than 25 characters")
	}

	return errs
}
