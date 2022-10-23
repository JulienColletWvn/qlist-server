package router

import (
	jwtauth "qlist/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	authMiddleware := jwtauth.New(jwtauth.Config{})

	auth := app.Group("/auth")
	Authentication(auth)

	api := app.Group("/api", authMiddleware)
	users := api.Group("/users")
	user := api.Group("/user")
	events := api.Group("/events")

	User(user)
	Users(users)
	Events(events)

}
