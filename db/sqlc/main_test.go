package db

import (
	"log"
	"os"
	"qlist/utils"

	"database/sql"
	"fmt"
	"testing"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	env := utils.GetEnvVariable
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"db", 5432, env("POSTGRES_USER"), env("POSTGRES_PASSWORD"), env("POSTGRES_DB"))

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(db)

	os.Exit(m.Run())
}
