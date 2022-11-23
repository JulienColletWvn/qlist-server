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

// GetUser godoc
// @Summary     Get current urser data
// @Description get user with current sessions
// @Tags        users
// @Produce     json
// @Param       auth header   int true "Authentication token"
// @Success     200  {object} User
// @Failure     401  {object} entities.HTTPError
// @Failure     404  {object} entities.HTTPError
// @Failure     500  {object} entities.HTTPError
// @Router      /users/{id} [get]
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
