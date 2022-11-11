package handler

import (
	"qlist/db/entities"
	jwtauth "qlist/middleware"
	"qlist/utils"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Username  string `json:"username"`
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
	var user entities.User
	userId, err := jwtauth.GetCurrentUserId(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	utils.Database.Where("id=?", userId).First(&user)

	return c.Status(fiber.StatusOK).JSON(User{
		Username:  user.Username,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Phone:     user.Phone,
	})
}
