package models

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// InitDB returns pointer of type DB
func InitDB(databaseSource string) {
	var err error
	db, err = sql.Open("sqlite3", databaseSource)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	_, err = db.Exec(` CREATE TABLE IF NOT EXISTS
		contact ( cont_id INTEGER PRIMARY KEY AUTOINCREMENT, 
		u_name TEXT NOT NULL, 
		pri_number CHAR(15), 
		address CHAR(70), 
		email_id CHAR(20))`)
	if err != nil {
		panic(err)
	}
}
