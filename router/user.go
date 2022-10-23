package router

import (
	handler "qlist/handler/user"

	"github.com/gofiber/fiber/v2"
)

func User(r fiber.Router) {
	r.Get("/", handler.GetUser)

	events := r.Group("/events")
	events.Post("/", handler.CreateEvent)
	events.Get("/", handler.GetUserEvents)
	events.Get("/:eventId", handler.GetUserEvent)
	events.Put("/:eventId", handler.UpdateUserEvent)
	events.Delete("/:eventId", handler.DeleteUserEvent)
}
