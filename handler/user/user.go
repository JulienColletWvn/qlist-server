package handler

import (
	"context"
	db "qlist/db/sqlc"
	jwtauth "qlist/middleware"
	"qlist/utils"

	"github.com/gofiber/fiber/v2"
)

type UserResponse struct {
	Username  string `json:"username"`
	CreatedAt string `json:"createdAt"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

func GetUser(c *fiber.Ctx) error {
	ctx := context.Background()
	userId, err := jwtauth.GetCurrentUserId(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	queries := db.New(utils.Database)

	u, err := queries.GetUserById(ctx, int32(userId))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(UserResponse{
		Firstname: u.Firstname,
		Lastname:  u.Lastname,
		Phone:     u.Phone,
		Email:     u.Email,
		Username:  u.Username,
		CreatedAt: u.CreatedAt.Time.UTC().String(),
	})
}
