package handler

import (
	"qlist/db/entities"
	jwtauth "qlist/middleware"
	"qlist/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateEvent(c *fiber.Ctx) error {
	event := new(entities.Event)
	user := new(entities.User)

	userId, err := jwtauth.GetCurrentUserId(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	if err := c.BodyParser(event); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	if err := utils.ValidateStruct(*event); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	utils.Database.Create(&event)
	utils.Database.Where("id=?", userId).First(&user)

	utils.Database.Model(&user).Association("Events").Append(event)

	return c.Status(fiber.StatusOK).JSON(event)
}

func GetUserEvent(c *fiber.Ctx) error {
	var event entities.Event
	eventId := c.AllParams()["eventId"]

	if eventId == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	user := new(entities.User)
	userId, err := jwtauth.GetCurrentUserId(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	utils.Database.Where("id=?", userId).First(&user)
	utils.Database.Model(&user).Where("id=?", eventId).Association("Events").Find(&event)

	if event.ID == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(event)
}

func GetUserEvents(c *fiber.Ctx) error {
	var user entities.User
	var events []entities.Event

	userId, err := jwtauth.GetCurrentUserId(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	utils.Database.Where("id=?", userId).First(&user)

	utils.Database.Model(&user).Association("Events").Find(&events)

	return c.Status(fiber.StatusOK).JSON(events)
}

func UpdateUserEvent(c *fiber.Ctx) error {
	var event entities.Event
	eventId := c.AllParams()["eventId"]

	eventUpdates := new(entities.Event)

	if eventId == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	user := new(entities.User)
	userId, err := jwtauth.GetCurrentUserId(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	if err := c.BodyParser(eventUpdates); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	utils.Database.Where("id=?", userId).First(&user)
	utils.Database.Model(&user).Where("id=?", eventId).Association("Events").Find(&event)

	if event.ID == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	utils.Database.Model(&event).Updates(eventUpdates)

	return c.Status(fiber.StatusOK).JSON(event)
}

func DeleteUserEvent(c *fiber.Ctx) error {
	var event entities.Event
	eventId := c.AllParams()["eventId"]

	if eventId == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	user := new(entities.User)
	userId, err := jwtauth.GetCurrentUserId(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	utils.Database.Where("id=?", userId).First(&user)
	utils.Database.Model(&user).Where("id=?", eventId).Association("Events").Find(&event)

	if event.ID == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	utils.Database.Delete(&event)

	return c.SendStatus(fiber.StatusNoContent)
}
