package controllers

import (
	"database/sql"
	"os"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("DB_URL"))
	// db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/kokutime?parseTime=true&loc=Asia%2Fjakarta")
	CheckError(err)
	return db
}
