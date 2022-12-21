package handler

import (
	"context"
	db "qlist/db/sqlc"
	jwtauth "qlist/middleware"
	"qlist/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CreateUserEventStewardBody struct {
	ContactId int `json:"contactId"`
}

func GetUserEventStewards(c *fiber.Ctx) error {
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

	guests, err := queries.GetUserEventStewards(ctx, db.GetUserEventStewardsParams{
		EventsID: int32(eventId),
		UsersID:  int32(userId),
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(guests)

}

func GetUserEventSteward(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)

	eventId, err := strconv.Atoi(c.AllParams()["eventId"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	stewardId, err := strconv.Atoi(c.AllParams()["stewardId"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userId, err := jwtauth.GetCurrentUserId(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	steward, err := queries.GetUserEventSteward(ctx, db.GetUserEventStewardParams{
		ID:       int32(stewardId),
		EventsID: int32(eventId),
		UsersID:  int32(userId),
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(steward)

}

func CreateUserEventSteward(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)
	body := CreateUserEventSellerBody{}

	eventId, err := strconv.Atoi(c.AllParams()["eventId"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userId, err := jwtauth.GetCurrentUserId(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	userEvents, err := queries.GetAdministratorEvents(ctx, int32(userId))

	if utils.GetIsUserOwningEvent(userEvents, eventId) == false {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	contacts, err := queries.GetUserContacts(ctx, int32(userId))

	if utils.GetIsUserOwningContact(contacts, body.ContactId) == false {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	_, creationError := queries.CreateUserEventSteward(ctx, db.CreateUserEventStewardParams{
		EventsID:   int32(eventId),
		ContactsID: int32(body.ContactId),
	})

	if creationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(creationError.Error())
	}

	return c.SendStatus(fiber.StatusCreated)
}

func DeleteUserEventSteward(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)

	sellerId, err := strconv.Atoi(c.AllParams()["stewardId"])

	userId, err := jwtauth.GetCurrentUserId(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	deletionError := queries.DeleteUserEventSteward(ctx, db.DeleteUserEventStewardParams{
		ID:      int32(sellerId),
		UsersID: int32(userId),
	})

	if deletionError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(deletionError.Error())
	}

	return c.SendStatus(fiber.StatusCreated)
}
