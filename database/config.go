package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Database struct {
	Db *sql.DB
}

func NewDatabase() (*Database, error) {
	connStr := "postgres://postgres:password123@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	return &Database{Db: db}, nil
}