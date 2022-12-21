package handler

import (
	"context"
	db "qlist/db/sqlc"
	jwtauth "qlist/middleware"
	"qlist/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Guest struct {
	ContactId int    `json:"contactId"`
	Note      string `json:"note"`
}

type GuestsCreationParams struct {
	Guests []Guest `json:"guests"`
}

func GetOwnedGuests(ownedContacts []db.Contact, guests []Guest) []Guest {
	contacts := []Guest{}

	for _, oc := range ownedContacts {
		for _, guest := range guests {
			if oc.ID == int32(guest.ContactId) {
				contacts = append(contacts, guest)
			}
		}
	}

	return contacts
}

func GetUserEventGuests(c *fiber.Ctx) error {
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

func GetUserEventGuest(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)

	eventId, err := strconv.Atoi(c.AllParams()["eventId"])
	guestId, err := strconv.Atoi(c.AllParams()["guestId"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userId, err := jwtauth.GetCurrentUserId(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	guest, err := queries.GetUserEventGuest(ctx, db.GetUserEventGuestParams{
		ID:       int32(guestId),
		EventsID: int32(eventId),
		UsersID:  int32(userId),
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(guest)

}

func CreateUserEventGuests(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)
	body := GuestsCreationParams{}
	creationParams := []db.CreateEventGuestsParams{}

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

	contacts, err := queries.GetUserContacts(ctx, int32(userId))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	ownedGuests := GetOwnedGuests(contacts, body.Guests)

	if len(ownedGuests) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	for _, guest := range ownedGuests {
		creationParams = append(creationParams, db.CreateEventGuestsParams{
			Note:       guest.Note,
			EventsID:   int32(eventId),
			ContactsID: int32(guest.ContactId),
		})
	}

	_, creationError := queries.CreateEventGuests(ctx, creationParams)

	if creationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(creationError.Error())
	}

	return c.SendStatus(fiber.StatusCreated)
}

func DeleteUserEventGuest(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)

	guestId, err := strconv.Atoi(c.AllParams()["guestId"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userId, err := jwtauth.GetCurrentUserId(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	deletetionError := queries.DeleteUserEventGuest(ctx, db.DeleteUserEventGuestParams{
		ID:      int32(guestId),
		UsersID: int32(userId),
	})

	if deletetionError != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusCreated)
}
