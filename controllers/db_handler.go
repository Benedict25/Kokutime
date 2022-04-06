package controllers

import (
	"database/sql"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/kokutime")
	// db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/kokutime?parseTime=true&loc=Asia%2Fjakarta")
	CheckError(err)
	return db
}
