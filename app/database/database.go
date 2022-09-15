package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func OpenDatabase() (db *sql.DB) {
	db, err := sql.Open("sqlite", "transactions.db")
	if err != nil {
		panic(err.Error())
	}

	log.Println("Opened database.")
	return
}
