package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var Database *sql.DB

func Connect() error {
	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable TimeZone=Europe/Brussels",
		"localhost", 5432, GetEnvVariable("POSTGRES_USER"), GetEnvVariable("POSTGRES_PASSWORD"), GetEnvVariable("POSTGRES_DB"))
	Database, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal("Connection with DB failed")
	}

	return nil
}
