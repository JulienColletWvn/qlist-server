package handler

import (
	"context"
	db "qlist/db/sqlc"
	jwtauth "qlist/middleware"
	"qlist/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CreateUserEventSellerBody struct {
	ContactId int `json:"contactId"`
}

func GetUserEventSellers(c *fiber.Ctx) error {
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

	sellers, err := queries.GetUserEventSellers(ctx, db.GetUserEventSellersParams{
		EventsID: int32(eventId),
		UsersID:  int32(userId),
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(sellers)

}

func GetUserEventSeller(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)

	eventId, err := strconv.Atoi(c.AllParams()["eventId"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	sellerId, err := strconv.Atoi(c.AllParams()["sellerId"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userId, err := jwtauth.GetCurrentUserId(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	sellers, err := queries.GetUserEventSeller(ctx, db.GetUserEventSellerParams{
		ID:       int32(sellerId),
		EventsID: int32(eventId),
		UsersID:  int32(userId),
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(sellers)

}

func CreateUserEventSeller(c *fiber.Ctx) error {
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

	_, creationError := queries.CreateUserEventSeller(ctx, db.CreateUserEventSellerParams{
		EventsID:   int32(eventId),
		ContactsID: int32(body.ContactId),
	})

	if creationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(creationError.Error())
	}

	return c.SendStatus(fiber.StatusCreated)
}

func DeleteUserEventSeller(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)

	sellerId, err := strconv.Atoi(c.AllParams()["sellerId"])

	userId, err := jwtauth.GetCurrentUserId(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	deletionError := queries.DeleteUserEventSeller(ctx, db.DeleteUserEventSellerParams{
		ID:      int32(sellerId),
		UsersID: int32(userId),
	})

	if deletionError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(deletionError.Error())
	}

	return c.SendStatus(fiber.StatusCreated)
}
