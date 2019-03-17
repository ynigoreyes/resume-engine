package models

// Comment defines the data necessary for representing a comment on a resume
type Comment struct {
	ResumeID    uint   `json:"resume_id" gorm:"not null"`
	CommenterID uint   `json:"commenter_id" gorm:"not null"`
	CommentBody string `json:"comment_body" gorm:"type:varchar(140);not null"`
}
