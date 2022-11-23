package router

import (
	handler "qlist/handler/events"
	user "qlist/handler/user"

	"github.com/gofiber/fiber/v2"
)

func User(r fiber.Router) {
	r.Get("/", user.GetUser)

	// contacts := r.Group("/contacts")
	// contacts.Post("/", handler.CreateContact)
	// contacts.Get("/", handler.GetUserContacts)

	events := r.Group("/events")
	events.Post("/", handler.CreateEvent)
	events.Get("/", handler.GetUserEvents)
	// events.Get("/:eventId", handler.GetUserEvent)
	// events.Put("/:eventId", handler.UpdateUserEvent)
	// events.Delete("/:eventId", handler.DeleteUserEvent)
	// events.Post("/:eventId/guests", handler.CreateGuest)
}
