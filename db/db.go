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

func CloseSQLiteDB(db *sql.DB) error {
	err := db.Close()
	if err != nil {
		return err
	}
	return nil
}
