package handler

import (
	"context"
	db "qlist/db/sqlc"
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

type User struct {
	Username  string `json:"username" validate:"required,min=3,max=32"`
	Email     string `json:"email" validate:"required,email"`
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	Phone     string `json:"phone"`
}

type CreateUserParams struct {
	User
	Password string `json:"password" validate:"required"`
}

func Register(c *fiber.Ctx) error {
	ctx := context.Background()
	user := CreateUserParams{}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	hash, err := hashPassword(user.Password)
	user.Password = hash

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})
	}

	if err := utils.ValidateStruct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	queries := db.New(utils.Database)

	u, err := queries.CreateUser(ctx, db.CreateUserParams{
		Username:  user.Username,
		Password:  hash,
		Phone:     user.Phone,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
	})

	if err != nil {
		userAlreayExists, _ := regexp.MatchString("duplicate key value violates unique constraint \"idx_users_email\"", err.Error())
		if userAlreayExists {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(utils.ApiError{
				Code: 1001,
				Text: "User already exists",
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	token, err := jwtauth.Encode(&jwt.StandardClaims{
		Subject:   strconv.Itoa(int(u.ID)),
		ExpiresAt: time.Now().UTC().Unix() + 24*60*100,
	})

	cookie := fiber.Cookie{
		Name:     "auth",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.Status(fiber.StatusOK).JSON(User{
		Username:  u.Username,
		Lastname:  u.Lastname,
		Firstname: u.Firstname,
		Email:     u.Email,
		Phone:     u.Phone,
	})

}
