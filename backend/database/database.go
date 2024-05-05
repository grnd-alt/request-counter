package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

var db *sql.DB

func Connect() {
	connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_URL"), os.Getenv("DB_NAME"))
	fmt.Println("CONNECTING NOW")
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("could not open", err)
	}
	pingerr := db.Ping()
	if pingerr != nil {
		fmt.Println("db not responding, retrying in 5 seconds")
		time.Sleep(5 * time.Second)
		for pingerr != nil {
			db, err = sql.Open("postgres", connStr)
			pingerr = db.Ping()
			if err != nil {
				log.Fatal("could not open", err)
			}
		}
	}
	ApplyMigrations()
}

func getDB() *sql.DB {
	return db
}
