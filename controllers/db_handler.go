package controllers

import (
	"database/sql"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/kokutime")
	// db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/kokutime?parseTime=true&loc=Asia%2Fjakarta")
	CheckError(err)
	return db
}

func connectGorm() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/kokutime?parseTime=true&loc=Asia%2FJakarta"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	return db
}
