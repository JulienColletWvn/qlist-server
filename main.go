package main

import (
	"log"
	"qlist/router"
	db "qlist/utils"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	db.Connect()
	app := fiber.New()
	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
