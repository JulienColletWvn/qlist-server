package handler

import (
	"context"
	db "qlist/db/sqlc"
	jwtauth "qlist/middleware"
	"qlist/utils"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	ctx := context.Background()
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	queries := db.New(utils.Database)

	user, err := queries.GetUserByEmail(ctx, data["email"])

	if user.ID == 0 || err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	token, err := jwtauth.Encode(&jwt.StandardClaims{
		Subject:   strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().UTC().Unix() + 24*60*100,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong",
		})
	}

	cookie := fiber.Cookie{
		Name:     "auth",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.SendStatus(fiber.StatusNoContent)

}
