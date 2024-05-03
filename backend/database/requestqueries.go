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

func GetRequests() []Request {
	db := getDB()
	rows, err := db.Query("select log_id,timestamp from request_logs")
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

func PutRequest() error {
	db := getDB()
	_, err := db.Exec("INSERT INTO request_logs (timestamp) VALUES (NOW())")
	return err
}
