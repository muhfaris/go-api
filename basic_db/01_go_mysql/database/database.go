package database

import (
	"database/sql"
	"fmt"
)

const (
	username = "dev"
	password = "dev123"
	dbname   = "db1"
	server   = "localhost"
	port     = 3306
)

func Connect() (*sql.DB, error) {
	dbtext := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, server, port, dbname)
	db, err := sql.Open("mysql", dbtext)

	if err == nil {

	}
	return db, nil
}
