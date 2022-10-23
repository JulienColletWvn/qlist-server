package router

import (
	handler "qlist/handler/users"

	"github.com/gofiber/fiber/v2"
)

func Users(r fiber.Router) {
	r.Get("/", handler.GetUsers)
}
