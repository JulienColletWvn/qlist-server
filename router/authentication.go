package router

import (
	handler "qlist/handler/authentication"

	"github.com/gofiber/fiber/v2"
)

func Authentication(r fiber.Router) {
	r.Post("/register", handler.Register)
	r.Post("/login", handler.Login)
	r.Get("/logout", handler.Logout)
}
