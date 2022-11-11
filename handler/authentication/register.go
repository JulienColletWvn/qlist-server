package handler

import (
	"qlist/db/entities"
	jwtauth "qlist/middleware"
	"qlist/utils"
	"regexp"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
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

	creationError := utils.Database.Create(&user).Error

	userAlreayExists, _ := regexp.MatchString("duplicate key value violates unique constraint \"users_email_key\"", creationError.Error())

	if userAlreayExists {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(utils.ApiError{
			Code: 1001,
			Text: "User already exists",
		})
	}

	token, err := jwtauth.Encode(&jwt.StandardClaims{
		Subject:   strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().UTC().Unix() + 24*60*100,
	})

	cookie := fiber.Cookie{
		Name:     "auth",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.Status(fiber.StatusOK).JSON(user)

}
