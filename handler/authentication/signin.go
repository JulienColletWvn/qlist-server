package handler

import (
	"qlist/db/entities"
	"qlist/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CreateUser(c *fiber.Ctx) error {
	user := new(entities.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	hash, err := hashPassword(user.Password)
	user.Password = hash

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})
	}

	if err := utils.ValidateStruct(*user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	utils.Database.Create(&user)

	return c.Status(fiber.StatusOK).JSON(user)

}
