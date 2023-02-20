package handler

import (
	"context"
	"errors"
	db "qlist/db/sqlc"
	jwtauth "qlist/middleware"
	"qlist/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"
)

type CreateTicketTransactionBody struct {
	StewardId int `json:"stewardId"`
	Amount    int `json:"amount"`
}

type UpdateTicketTransactionStatusBody struct {
	Status db.TransactionStatus `json:"status"`
}

func GetIsOwningGuest(c *fiber.Ctx, userId int, guestId int, eventId int) (bool, error) {
	ctx := context.Background()
	queries := db.New(utils.Database)

	guests, err := queries.GetUserEventGuests(ctx, db.GetUserEventGuestsParams{
		EventsID: int32(eventId),
		UsersID:  int32(userId),
	})

	if err != nil {
		return false, errors.New("Can't get user event guests")
	}

	return slices.IndexFunc(guests, func(g db.Guest) bool { return g.ID == int32(guestId) }) != -1, nil
}

func GetGuestTicketTransactions(c *fiber.Ctx) error {
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

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	userIsOwningGuest, err := GetIsOwningGuest(c, userId, guestId, eventId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	if userIsOwningGuest != true {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	transactions, err := queries.GetGuestTicketTransactions(ctx, int32(ticketId))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(transactions)

}

func GetGuestTicketTransaction(c *fiber.Ctx) error {
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

	transactionId, err := strconv.Atoi(c.AllParams()["transactionId"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userId, err := jwtauth.GetCurrentUserId(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	userIsOwningGuest, err := GetIsOwningGuest(c, userId, guestId, eventId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	if userIsOwningGuest != true {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	transaction, err := queries.GetGuestTicketTransaction(ctx, db.GetGuestTicketTransactionParams{
		TicketsID: int32(ticketId),
		ID:        int32(transactionId),
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(transaction)

}

func CreateGuestTicketTransaction(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)
	ticketTransaction := CreateTicketTransactionBody{}

	if err := c.BodyParser(&ticketTransaction); err != nil {
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

	ticketId, err := strconv.Atoi(c.AllParams()["ticketId"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userId, err := jwtauth.GetCurrentUserId(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	userIsOwningGuest, err := GetIsOwningGuest(c, userId, guestId, eventId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	if userIsOwningGuest != true {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	transaction, err := queries.CreateGuestTicketTransaction(ctx, db.CreateGuestTicketTransactionParams{
		TicketsID:  int32(ticketId),
		StewardsID: int32(ticketTransaction.StewardId),
		Amount:     int32(ticketTransaction.Amount),
		Status:     db.TransactionStatusPending,
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(transaction)

}

func UpdateGuestTicketTransactionStatus(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)
	ticketTransactionStatusBody := UpdateTicketTransactionStatusBody{}

	if err := c.BodyParser(&ticketTransactionStatusBody); err != nil {
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

	transactionId, err := strconv.Atoi(c.AllParams()["transactionId"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userId, err := jwtauth.GetCurrentUserId(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	userIsOwningGuest, err := GetIsOwningGuest(c, userId, guestId, eventId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	if userIsOwningGuest != true {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	transaction, err := queries.UpdateGuestTicketTransactionStatus(ctx, db.UpdateGuestTicketTransactionStatusParams{
		Status: ticketTransactionStatusBody.Status,
		ID:     int32(transactionId),
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(transaction)
}
