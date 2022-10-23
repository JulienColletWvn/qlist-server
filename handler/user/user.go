package handler

import (
	"qlist/db/entities"
	jwtauth "qlist/middleware"
	"qlist/utils"

	"github.com/gofiber/fiber/v2"
)

// GetUser godoc
// @Summary     Get current urser data
// @Description get user with current sessions
// @Tags        users
// @Produce     json
// @Param       auth header   int true "Authentication token"
// @Success     200  {object} entities.User
// @Failure     401  {object} entities.HTTPError
// @Failure     404  {object} entities.HTTPError
// @Failure     500  {object} entities.HTTPError
// @Router      /users/{id} [get]
func GetUser(c *fiber.Ctx) error {
	var user entities.User
	userId, err := jwtauth.GetCurrentUserId(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	utils.Database.Where("id=?", userId).First(&user)

	return c.Status(fiber.StatusOK).JSON(user)
}
