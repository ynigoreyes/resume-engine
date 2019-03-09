package api

import (
	"fmt"
	"github.com/jinzhu/gorm"
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

// CommentHandler exposes the endpoint necessary for handling resume comments
func (*Routes) CommentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "AAAAA")
}
