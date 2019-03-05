package models

import (
	"github.com/jinzhu/gorm"
)

// Comment defines the data necessary for representing a comment on a resume
type Comment struct {
	gorm.Model
	CommenterID string `json:"commenter_id"`
	ResumeID    string `json:"resume_id"`
	CommentBody string `json:"comment_body" gorm:"type:varchar(140)"`
}
