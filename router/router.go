package router

import (
	jwtauth "qlist/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	authMiddleware := jwtauth.New(jwtauth.Config{})
	api := app.Group("/api")

	auth := api.Group("/auth")
	Authentication(auth)

	user := api.Group("/user", authMiddleware)

	User(user)

}
