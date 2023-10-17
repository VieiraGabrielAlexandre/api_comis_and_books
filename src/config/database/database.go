package database

import (
	"database/sql"
	"fmt"
	"os"
)

var Db *sql.DB

func Connect() {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?parseTime=true"

	var err error

	Db, err = sql.Open("mysql", dsn)

	if err != nil {
		panic(err.Error())
	}

	err = Db.Ping()

	if err != nil {
		panic(err.Error())
	}

	defer Db.Close()
	fmt.Println("Successfully connected to database")
}
