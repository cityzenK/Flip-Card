package main

import (
	"cards/cmd/api"
	"cards/db"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func initCards(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB: Successfully connected!")
}

func main() {
	db, err := db.NewSQLite("./db/cards-jwasham-extreme(1).db")
	if err != nil {
		log.Fatal(err)
	}

	initCards(db)

	server := api.NewAPIServer(":8088", db)
	if err := server.Run(); err != nil {
		log.Fatal()
	}
}
