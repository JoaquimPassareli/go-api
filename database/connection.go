package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./pessoas.db")
	if err != nil {
		panic(err)
	}
	return db
}
