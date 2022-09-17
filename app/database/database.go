package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func OpenDatabase() (db *sql.DB) {
	db, err := sql.Open("sqlite", "transactions.db")
	if err != nil {
		panic(err.Error())
	}
	DB = db
	log.Println("Opened database.")
	return
}
