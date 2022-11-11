package handler

import (
	"qlist/db/entities"
	jwtauth "qlist/middleware"
	"qlist/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateContact(c *fiber.Ctx) error {
	contact := entities.Contact{}
	contacts := []entities.Contact{}
	user := new(entities.User)

	userId, err := jwtauth.GetCurrentUserId(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	parsingErr := c.BodyParser(&contact)
	parsingMultipleError := c.BodyParser(&contacts)

	if parsingErr != nil || parsingMultipleError != nil {
		c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	if parsingMultipleError == nil {
		errors := [][]*utils.ErrorResponse{}
		for _, contact := range contacts {
			if err := utils.ValidateStruct(&contact); err != nil {
				errors = append(errors, err)
			}
		}

		if len(errors) > 0 {
			return c.Status(fiber.StatusBadRequest).JSON(errors)
		}

		utils.Database.Create(&contacts)
		utils.Database.Where("id=?", userId).First(&user)

		utils.Database.Model(&user).Association("Contacts").Append(&contacts)

		return c.Status(fiber.StatusOK).JSON(contacts)

	}

	if err := utils.ValidateStruct(&contact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	utils.Database.Create(&contact)
	utils.Database.Where("id=?", userId).First(&user)

	utils.Database.Model(&user).Association("Contacts").Append(&contact)

	return c.Status(fiber.StatusOK).JSON(contact)
}

func GetUserContacts(c *fiber.Ctx) error {
	var user entities.User
	var contacts []entities.Contact

	userId, err := jwtauth.GetCurrentUserId(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	utils.Database.Where("id=?", userId).First(&user)

	utils.Database.Model(&user).Association("Contacts").Find(&contacts)

	return c.Status(fiber.StatusOK).JSON(contacts)

}

func DeleteUserContact(c *fiber.Ctx) error {
	var contact entities.Contact
	contactId := c.AllParams()["contactId"]

	if contactId == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	user := new(entities.User)
	userId, err := jwtauth.GetCurrentUserId(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	utils.Database.Where("id=?", userId).First(&user)
	utils.Database.Model(&user).Where("id=?", contactId).Association("Contacts").Find(&contact)

	if contact.ID == 0 {
		return c.SendStatus(fiber.StatusNotFound)
	}

	utils.Database.Delete(&contact)

	return c.SendStatus(fiber.StatusNoContent)
}
