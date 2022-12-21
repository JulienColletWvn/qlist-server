package handler

import (
	"context"
	"database/sql"
	db "qlist/db/sqlc"
	jwtauth "qlist/middleware"
	"qlist/utils"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type CreateUserEventTicketTypeBody struct {
	EventId           int    `json:"eventId"`
	Name              string `json:"name"`
	StartValidityDate string `json:"startValidityDate"`
	EndValidityDate   string `json:"endValidityDate"`
	UsageLimitation   int    `json:"usageLimitation"`
	UsageUnimited     bool   `json:"usageUnlimited"`
}

func GetUserEventTicketTypes(c *fiber.Ctx) error {
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

	ticketsTypes, err := queries.GetEventTicketsTypes(ctx, db.GetEventTicketsTypesParams{
		EventsID: int32(eventId),
		UsersID:  int32(userId),
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(ticketsTypes)
}

func GetUserEventTicketType(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)

	ticketTypeId, err := strconv.Atoi(c.AllParams()["ticketTypeId"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userId, err := jwtauth.GetCurrentUserId(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	ticketsType, err := queries.GetEventTicketsType(ctx, db.GetEventTicketsTypeParams{
		ID:      int32(ticketTypeId),
		UsersID: int32(userId),
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(ticketsType)
}

func CreateUserEventTicketType(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)
	body := CreateUserEventTicketTypeBody{}

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

	ticketType, creationError := queries.CreateEventTicketType(ctx, db.CreateEventTicketTypeParams{
		EventsID: int32(eventId),
		Name:     body.Name,
		StartValidityDate: sql.NullTime{
			Time:  startDate,
			Valid: true,
		},
		EndValidityDate: sql.NullTime{
			Time:  endDate,
			Valid: true,
		},
		UsageLimitation: int32(body.UsageLimitation),
		UsageUnlimited:  body.UsageUnimited,
	})

	if creationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(creationError.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(ticketType)
}

func DeleteUserEventTicketType(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)

	ticketTypeId, err := strconv.Atoi(c.AllParams()["ticketTypeId"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userId, err := jwtauth.GetCurrentUserId(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	deletionError := queries.DeleteEventTicketsType(ctx, db.DeleteEventTicketsTypeParams{
		ID:      int32(ticketTypeId),
		UsersID: int32(userId),
	})

	if deletionError != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusCreated)
}
