package handler

import (
	"qlist/db/entities"
	"qlist/utils"

	"github.com/gofiber/fiber/v2"
)

func GetEvent(c *fiber.Ctx) error {
	var users []entities.User

	res := utils.Database.Find(&users)

	if res.Error != nil {
		c.Status(fiber.StatusInternalServerError).SendString(res.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(users)
}
