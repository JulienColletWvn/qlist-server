package handler

import (
	"context"
	db "qlist/db/sqlc"
	jwtauth "qlist/middleware"
	"qlist/utils"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type CreateUserEventWalletTypeBody struct {
	EventId           int    `json:"eventId"`
	Name              string `json:"name"`
	StartValidityDate string `json:"startValidityDate"`
	EndValidityDate   string `json:"endValidityDate"`
	MaxAmount         int    `json:"maxAmount"`
	OnelineReload     bool   `json:"onlineReload"`
}

func GetUserEventWalletTypes(c *fiber.Ctx) error {
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

	walletsTypes, err := queries.GetEventWalletsTypes(ctx, db.GetEventWalletsTypesParams{
		EventsID: int32(eventId),
		UsersID:  int32(userId),
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(walletsTypes)
}

func GetUserEventWalletType(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)

	walletTypeId, err := strconv.Atoi(c.AllParams()["walletTypeId"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userId, err := jwtauth.GetCurrentUserId(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	walletType, err := queries.GetEventWalletsType(ctx, db.GetEventWalletsTypeParams{
		ID:      int32(walletTypeId),
		UsersID: int32(userId),
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(walletType)
}

func CreateUserEventWalletType(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)
	body := CreateUserEventWalletTypeBody{}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	eventId, err := strconv.Atoi(c.AllParams()["eventId"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userId, err := jwtauth.GetCurrentUserId(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	userEvents, err := queries.GetAdministratorEvents(ctx, int32(userId))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	if utils.GetIsUserOwningEvent(userEvents, eventId) == false {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	startDate, err := time.Parse(time.RFC3339, body.StartValidityDate)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	endDate, err := time.Parse(time.RFC3339, body.EndValidityDate)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	walletType, creationError := queries.CreateEventWalletType(ctx, db.CreateEventWalletTypeParams{
		EventsID:          int32(eventId),
		Name:              body.Name,
		StartValidityDate: startDate,
		EndValidityDate:   endDate,
		MaxAmount:         int32(body.MaxAmount),
		OnlineReload:      body.OnelineReload,
	})

	if creationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(creationError.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(walletType)
}

func DeleteUserEventWalletType(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)

	walletTypeId, err := strconv.Atoi(c.AllParams()["walletTypeId"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userId, err := jwtauth.GetCurrentUserId(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	deletionError := queries.DeleteEventWalletType(ctx, db.DeleteEventWalletTypeParams{
		ID:      int32(walletTypeId),
		UsersID: int32(userId),
	})

	if deletionError != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusCreated)
}
