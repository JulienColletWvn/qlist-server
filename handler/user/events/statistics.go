package handler

import (
	"context"
	db "qlist/db/sqlc"
	jwtauth "qlist/middleware"
	"qlist/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetUserEventStatistics(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)

	eventId, err := strconv.Atoi(c.AllParams()["eventId"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userId, err := jwtauth.GetCurrentUserId(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	guests, err := queries.GetUserEventGuests(ctx, db.GetUserEventGuestsParams{
		EventsID: int32(eventId),
		UsersID:  int32(userId),
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(guests)

}
