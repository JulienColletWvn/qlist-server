package handler

import (
	"qlist/db/entities"
	jwtauth "qlist/middleware"
	"qlist/utils"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var users []entities.User

	res := utils.Database.Find(&users)

	if res.Error != nil {
		c.Status(fiber.StatusInternalServerError).SendString(res.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	var user entities.User
	userId, err := jwtauth.GetCurrentUserId(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	utils.Database.Where("id=?", userId).First(&user)

	return c.Status(fiber.StatusOK).JSON(user)
}
