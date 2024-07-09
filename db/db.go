package db

import (
	"database/sql"
	"log"
)

func NewSQLite(cfg string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", cfg)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

