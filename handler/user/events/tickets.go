package handler

import (
	"context"
	db "qlist/db/sqlc"
	jwtauth "qlist/middleware"
	"qlist/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CreateGuestTicketBody struct {
	GuestID      int `json:"guestId"`
	TicketTypeID int `json:"ticketTypeId"`
}

func hasUserRightsOnGuest(c *fiber.Ctx, userId int, eventId int, guestId int) (bool, error) {
	ctx := context.Background()
	queries := db.New(utils.Database)

	guest, err := queries.GetUserEventGuest(ctx, db.GetUserEventGuestParams{
		ID:       int32(guestId),
		EventsID: int32(eventId),
		UsersID:  int32(userId),
	})

	if err != nil {
		return false, c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userEvents, err := queries.GetAdministratorEvents(ctx, int32(userId))

	if utils.GetIsUserOwningEvent(userEvents, int(guest.EventsID)) == false {
		return false, c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	return true, nil
}

func GetGuestTicket(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)

	eventId, err := strconv.Atoi(c.AllParams()["eventId"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	guestId, err := strconv.Atoi(c.AllParams()["guestId"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	ticketId, err := strconv.Atoi(c.AllParams()["ticketId"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userId, err := jwtauth.GetCurrentUserId(c)
	hasRight, hasRightError := hasUserRightsOnGuest(c, userId, eventId, guestId)

	if err != nil || hasRightError != nil || hasRight == false {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	guestTicket, err := queries.GetGuestTicket(ctx, db.GetGuestTicketParams{
		GuestsID: int32(guestId),
		ID:       int32(ticketId),
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(guestTicket)
}

func GetGuestTickets(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)

	eventId, err := strconv.Atoi(c.AllParams()["eventId"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	guestId, err := strconv.Atoi(c.AllParams()["guestId"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userId, err := jwtauth.GetCurrentUserId(c)
	hasRight, hasRightError := hasUserRightsOnGuest(c, userId, eventId, guestId)

	if err != nil || hasRightError != nil || hasRight == false {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	guestTickets, err := queries.GetGuestTickets(ctx, int32(guestId))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(guestTickets)
}

func CreateGuestTicket(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)
	body := CreateGuestTicketBody{}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	eventId, err := strconv.Atoi(c.AllParams()["eventId"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	guestId, err := strconv.Atoi(c.AllParams()["guestId"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userId, err := jwtauth.GetCurrentUserId(c)
	hasRight, hasRightError := hasUserRightsOnGuest(c, userId, eventId, guestId)

	if err != nil || hasRightError != nil || hasRight == false {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	ticket, creationError := queries.CreateGuestTicket(ctx, db.CreateGuestTicketParams{
		TicketsTypeID: int32(body.TicketTypeID),
		GuestsID:      int32(body.GuestID),
	})

	if creationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(creationError.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(ticket)
}

func DeleteGuestTicket(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)

	eventId, err := strconv.Atoi(c.AllParams()["eventId"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	ticketId, err := strconv.Atoi(c.AllParams()["ticketId"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	guestId, err := strconv.Atoi(c.AllParams()["guestId"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userId, err := jwtauth.GetCurrentUserId(c)
	hasRight, hasRightError := hasUserRightsOnGuest(c, userId, eventId, guestId)

	if err != nil || hasRightError != nil || hasRight == false {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	deletionError := queries.DeleteGuestTicket(ctx, int32(ticketId))

	if deletionError != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusCreated)
}
