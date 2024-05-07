package database

import (
	"fmt"
	"log"
	"time"
)

type Request struct {
	Log_id    uint
	Timestamp time.Time
}

func GetRequests(access_key string) []Request {
	db := getDB()
	rows, err := db.Query("select log_id,timestamp from request_logs where access_key = $1", access_key)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var requests []Request
	for rows.Next() {
		var request Request

		err := rows.Scan(&request.Log_id, &request.Timestamp)
		if err != nil {
			fmt.Println("ERROR SCANNING LINE", err)
		}

		requests = append(requests, request)
	}

	return requests
}

func GetRequestsCount(access_key string) int {
	db := getDB()
	rows, err := db.Query("select count(log_id) from request_logs where access_key = $1", access_key)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	rows.Next()
	var res int
	rows.Scan(&res)
	return res
}

func PutRequest(access_key string) error {
	db := getDB()
	_, err := db.Exec("INSERT INTO request_logs (timestamp,access_key) VALUES (NOW(),$1)", access_key)
	return err
}
