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
	user := api.Group("/user")
	users := api.Group("/users")
	events := api.Group("/events")

	User(user)
	Users(users)
	Events(events)

}
