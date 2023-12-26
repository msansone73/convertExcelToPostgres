package model

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func getConectionString() string {
	return "user=msansone dbname=sansone password=senha host=149.102.140.240 port=5432 sslmode=disable"
}

func getConnection() *sql.DB {
	db, err := sql.Open("postgres", getConectionString())
	if err!=nil {
		log.Fatal(err.Error())
	}

	return db
}