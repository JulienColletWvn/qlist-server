package router

import (
	handler "qlist/handler/authentication"

	"github.com/gofiber/fiber/v2"
)

func Authentication(r fiber.Router) {
	r.Post("/register", handler.CreateUser)
	r.Post("/login", handler.GetUser)
	r.Get("/logout", handler.Logout)
}
