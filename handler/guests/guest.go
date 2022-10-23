package handler

import (
	"qlist/db/entities"
	"qlist/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateGuest(c *fiber.Ctx) error {
	event := new(entities.Event)

	if err := c.BodyParser(event); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	if err := utils.ValidateStruct(*event); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	utils.Database.Create(&event)

	return c.Status(fiber.StatusOK).JSON(event)

}

func UpdateGuest(c *fiber.Ctx) error {
	var users []entities.User

	res := utils.Database.Find(&users)

	if res.Error != nil {
		c.Status(fiber.StatusInternalServerError).SendString(res.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func DeleteGuest(c *fiber.Ctx) error {
	var users []entities.User

	res := utils.Database.Find(&users)

	if res.Error != nil {
		c.Status(fiber.StatusInternalServerError).SendString(res.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func GetGuest(c *fiber.Ctx) error {
	var users []entities.User

	res := utils.Database.Find(&users)

	if res.Error != nil {
		c.Status(fiber.StatusInternalServerError).SendString(res.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(users)
}
