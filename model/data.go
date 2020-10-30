package model

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func getDBClient() (client *sql.DB) {

	client, err := sql.Open("sqlite3", "./cherry.sqlite3")

	if err != nil {
		log.Fatalf("error connecting to database: %v\n", err)
	}

	return client
}
