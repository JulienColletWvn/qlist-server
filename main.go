package main

import (
	"log"
	"qlist/router"
	"qlist/utils"
	db "qlist/utils"

	_ "qlist/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	_ "github.com/lib/pq"
)

func main() {
	db.Connect()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     utils.GetEnvVariable("CORS_ORIGINS"),
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	app.Get("/swagger/*", swagger.HandlerDefault)

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
