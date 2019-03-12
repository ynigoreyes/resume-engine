package models

import (
	"github.com/jinzhu/gorm"
	"net/url"
)

// Comment defines the data necessary for representing a comment on a resume
type Comment struct {
	gorm.Model
	CommenterID string `json:"commenter_id" gorm:"foreignkey:UserID"`
	CommentBody string `json:"comment_body" gorm:"type:varchar(140)"`
}

// Validate valides a comment creation request
func (c *Comment) Validate() url.Values {
	// Track request body errors for response
	errs := url.Values{}

	// Check if the commenter's id empty
	if c.CommenterID == "" {
		errs.Add("commenter_id", "commenter_id field is required")
	}

	// Check if the comment body is empty
	if c.CommentBody == "" {
		errs.Add("comment_body", "comment_body field is required")
	}

	return errs
}
