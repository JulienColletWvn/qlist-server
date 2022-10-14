package handler

import (
	"database/sql"
	db "qlist/db/sqlc"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type Handler func(c *fiber.Ctx) error

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func MakeCreateUser(database *sql.DB) Handler {
	return func(c *fiber.Ctx) error {
		queries := db.New(database)
		u := new(db.CreateUserParams)

		if err := c.BodyParser(&u); err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})

		}

		hash, err := hashPassword(u.Password)
		u.Password = hash

		if err != nil {
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})

		}

		user, error := queries.CreateUser(c.Context(), *u)

		if error != nil {
			c.SendString(error.Error())
		}

		return c.JSON(user)

	}
}

func GetUsers(c *fiber.Ctx) error {
	return c.SendString("about")
}
