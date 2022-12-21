package handler

import (
	"context"
	"database/sql"
	db "qlist/db/sqlc"
	jwtauth "qlist/middleware"
	"qlist/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Contact struct {
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Phone     string `json:"phone"`
	Lang      string `json:"lang"`
}

type ContactError struct {
	Email string `json:"email"`
	Error string `json:"error"`
}

func CreateContacts(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)
	contacts := []Contact{}
	contactsParams := []db.CreateUserContactsParams{}
	userId, err := jwtauth.GetCurrentUserId(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	if err := c.BodyParser(&contacts); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	for _, contact := range contacts {
		contactsParams = append(contactsParams, db.CreateUserContactsParams{
			Email:     contact.Email,
			Firstname: contact.Firstname,
			Lastname:  contact.Lastname,
			Phone:     contact.Phone,
			Lang: sql.NullString{
				String: contact.Lang,
				Valid:  true,
			},
			CreatorID: int32(userId),
		})
	}

	_, creationError := queries.CreateUserContacts(ctx, contactsParams)

	if creationError != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusOK)
}

func GetUserContacts(c *fiber.Ctx) error {
	ctx := context.Background()
	userId, err := jwtauth.GetCurrentUserId(c)
	queries := db.New(utils.Database)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	contacts, err := queries.GetUserContacts(ctx, int32(userId))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(contacts)

}

func GetUserContact(c *fiber.Ctx) error {
	ctx := context.Background()
	userId, err := jwtauth.GetCurrentUserId(c)
	queries := db.New(utils.Database)

	contactId, err := strconv.Atoi(c.AllParams()["contactId"])

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	contact, err := queries.GetUserContact(ctx, db.GetUserContactParams{
		CreatorID: int32(userId),
		ID:        int32(contactId),
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(contact)
}

func DeleteUserContact(c *fiber.Ctx) error {
	ctx := context.Background()
	userId, err := jwtauth.GetCurrentUserId(c)
	queries := db.New(utils.Database)

	contactId, err := strconv.Atoi(c.AllParams()["contactId"])

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	deletionError := queries.DeleteUserContact(ctx, db.DeleteUserContactParams{
		CreatorID: int32(userId),
		ID:        int32(contactId),
	})

	if deletionError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(deletionError.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}
