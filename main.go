package main

import (
	"log"
	"qlist/router"
	db "qlist/utils"

	_ "qlist/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/lib/pq"
)

func main() {
	db.Connect()

	app := fiber.New()
	app.Get("/swagger/*", swagger.HandlerDefault)

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
