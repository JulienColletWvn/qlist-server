package router

import (
	"database/sql"
	"qlist/handler"

	"github.com/gofiber/fiber/v2"
)

func Users(r fiber.Router, db *sql.DB) {
	r.Post("/", handler.MakeCreateUser(db))
	r.Get("/:userId", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
