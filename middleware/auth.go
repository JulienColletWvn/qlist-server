package jwtauth

import (
	"errors"
	"qlist/utils"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type Config struct {
	Filter       func(c *fiber.Ctx) bool
	Unauthorized fiber.Handler
	Decode       func(c *fiber.Ctx) (*jwt.StandardClaims, error)
	Expiry       int64
}

var ConfigDefault = Config{
	Filter:       nil,
	Decode:       nil,
	Unauthorized: nil,
	Expiry:       60,
}

func configDefault(config ...Config) Config {
	if len(config) < 1 {
		return ConfigDefault
	}

	cfg := config[0]

	if cfg.Filter == nil {
		cfg.Filter = ConfigDefault.Filter
	}

	if cfg.Expiry == 0 {
		cfg.Expiry = ConfigDefault.Expiry
	}

	if cfg.Decode == nil {
		cfg.Decode = func(c *fiber.Ctx) (*jwt.StandardClaims, error) {

			authCookie := c.Cookies("auth")

			token, err := jwt.ParseWithClaims(authCookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
				return []byte(utils.GetEnvVariable("OAUTH_SECRET_KEY")), nil
			})

			if err != nil {
				return nil, errors.New("Error parsing token")
			}

			claims, ok := token.Claims.(*jwt.StandardClaims)

			if !(ok && token.Valid) {
				return nil, errors.New("Invalid token")
			}

			if claims.ExpiresAt < time.Now().UTC().Unix() {
				return nil, errors.New("jwt is expired")
			}

			return claims, nil
		}
	}

	if cfg.Unauthorized == nil {
		cfg.Unauthorized = func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
	}

	return cfg
}

func Encode(claims *jwt.StandardClaims) (string, error) {
	c := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := c.SignedString([]byte(utils.GetEnvVariable("OAUTH_SECRET_KEY")))

	if err != nil {
		return "", errors.New("Error creating a token")
	}

	return token, nil
}

func New(config Config) fiber.Handler {
	cfg := configDefault(config)

	return func(c *fiber.Ctx) error {
		if cfg.Filter != nil && cfg.Filter(c) {
			return c.Next()
		}

		claims, err := cfg.Decode(c)

		if err == nil {
			c.Locals("authClaims", claims)
			return c.Next()
		}

		return cfg.Unauthorized(c)
	}
}

func GetCurrentClaims(c *fiber.Ctx) (jwt.StandardClaims, error) {
	l := c.Locals("authClaims")

	claims, ok := l.(*jwt.StandardClaims)

	if !(ok) {
		return jwt.StandardClaims{}, errors.New("No claim")
	}

	return *claims, nil

}

func GetCurrentUserId(c *fiber.Ctx) (int, error) {
	claims, err := GetCurrentClaims(c)

	if err != nil || claims.Subject == "" {
		return 0, err
	}

	userId, err := strconv.Atoi(claims.Subject)

	if err != nil {
		return 0, errors.New("Wrong user id")
	}

	return userId, nil
}
