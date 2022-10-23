package router

import (
	handler "qlist/handler/users"

	"github.com/gofiber/fiber/v2"
)

func User(r fiber.Router) {
	events := r.Group("/events")
	events.Post("/", handler.CreateEvent)
	events.Get("/", handler.GetUserEvents)
	events.Get("/:eventId", handler.GetUserEvent)
	events.Put("/:eventId", handler.UpdateUserEvent)
	events.Delete("/:eventId", handler.DeleteUserEvent)
}
