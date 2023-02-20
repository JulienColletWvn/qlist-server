package utils

import (
	"context"
	db "qlist/db/sqlc"

	"github.com/gofiber/fiber/v2"
)

func HasUserRightsOnGuest(c *fiber.Ctx, userId int, eventId int, guestId int) (bool, error) {
	ctx := context.Background()
	queries := db.New(Database)

	guest, err := queries.GetUserEventGuest(ctx, db.GetUserEventGuestParams{
		ID:       int32(guestId),
		EventsID: int32(eventId),
		UsersID:  int32(userId),
	})

	if err != nil {
		return false, c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userEvents, err := queries.GetAdministratorEvents(ctx, int32(userId))

	if GetIsUserOwningEvent(userEvents, int(guest.EventsID)) == false {
		return false, c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	return true, nil
}
