package router

import (
	handler "qlist/handler/events"

	"github.com/gofiber/fiber/v2"
)

func Events(r fiber.Router) {
	r.Get("/:eventId", handler.GetEvent)
}
