package handler

import (
	"context"
	db "qlist/db/sqlc"
	jwtauth "qlist/middleware"
	"qlist/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CreateUserEventCashierBody struct {
	ContactId int `json:"contactId"`
}

func CreateUserEventCashier(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)
	body := CreateUserEventCashierBody{}

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

	_, creationError := queries.CreateUserEventCashier(ctx, db.CreateUserEventCashierParams{
		EventsID:   int32(eventId),
		ContactsID: int32(body.ContactId),
	})

	if creationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(creationError.Error())
	}

	return c.SendStatus(fiber.StatusCreated)
}

func DeleteUserEventCashier(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)

	cashierId, err := strconv.Atoi(c.AllParams()["cashierId"])

	userId, err := jwtauth.GetCurrentUserId(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	deletionError := queries.DeleteUserEventCashier(ctx, db.DeleteUserEventCashierParams{
		ID:      int32(cashierId),
		UsersID: int32(userId),
	})

	if deletionError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(deletionError.Error())
	}

	return c.SendStatus(fiber.StatusCreated)
}

func GetUserEventCashiers(c *fiber.Ctx) error {
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

	guests, err := queries.GetUserEventCashiers(ctx, db.GetUserEventCashiersParams{
		EventsID: int32(eventId),
		UsersID:  int32(userId),
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(guests)

}

func GetUserEventCashier(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)

	eventId, err := strconv.Atoi(c.AllParams()["eventId"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	cashierId, err := strconv.Atoi(c.AllParams()["cashierId"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userId, err := jwtauth.GetCurrentUserId(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	guests, err := queries.GetUserEventCashier(ctx, db.GetUserEventCashierParams{
		ID:       int32(cashierId),
		EventsID: int32(eventId),
		UsersID:  int32(userId),
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(guests)

}
