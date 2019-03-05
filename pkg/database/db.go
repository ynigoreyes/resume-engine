package database

import (
	"database/sql"
	// "github.com/go-sql-driver/mysql"
)

// Init initializes an instance of the database
func Init() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root@cloudsql(project-id:instance-name)/dbname")
	if err != nil {
		return nil, err
	}
	return db, nil
}
