package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var db *sql.DB

func Connect() {
	connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_URL"), os.Getenv("DB_NAME"))
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	ApplyMigrations()
}

func getDB() *sql.DB {
	return db
}
