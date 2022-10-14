package main

import (
	"database/sql"
	"fmt"
	"log"
	"qlist/router"
	"qlist/utils"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	env := utils.GetEnvVariable
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"db", 5432, env("POSTGRES_USER"), env("POSTGRES_PASSWORD"), env("POSTGRES_DB"))
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal("Connection with DB failed")
	}

	app := fiber.New()

	router.SetupRoutes(app, db)

	log.Fatal(app.Listen(":3000"))
}
