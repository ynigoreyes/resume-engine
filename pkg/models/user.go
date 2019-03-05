package models

// User defines a user interacting with the application
type User struct {
	UserID    string `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	TagLine   string `json:"tag_line"`
	ResumeID  string `json:"resume_id"`
}
