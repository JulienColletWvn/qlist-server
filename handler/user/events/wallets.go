package handler

import (
	"context"
	"database/sql"
	db "qlist/db/sqlc"
	jwtauth "qlist/middleware"
	"qlist/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CreateGuestWalletBody struct {
	GuestID      int  `json:"guestId"`
	WalletTypeID int  `json:"walletTypeId"`
	MaxAmount    int  `json:"maxAmount"`
	OnlineReload bool `json:"onlineReload"`
}

func GetGuestWallet(c *fiber.Ctx) error {
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

	walletId, err := strconv.Atoi(c.AllParams()["walletId"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userId, err := jwtauth.GetCurrentUserId(c)
	hasRight, hasRightError := utils.HasUserRightsOnGuest(c, userId, eventId, guestId)

	if err != nil || hasRightError != nil || hasRight == false {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	guestWallet, err := queries.GetGuestWallet(ctx, db.GetGuestWalletParams{
		GuestsID: int32(guestId),
		ID:       int32(walletId),
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(guestWallet)
}

func GetGuestWallets(c *fiber.Ctx) error {
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
	hasRight, hasRightError := utils.HasUserRightsOnGuest(c, userId, eventId, guestId)

	if err != nil || hasRightError != nil || hasRight == false {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	guestWallets, err := queries.GetGuestWallets(ctx, int32(guestId))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(guestWallets)
}

func CreateGuestWallet(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)
	body := CreateGuestWalletBody{}

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
	hasRight, hasRightError := utils.HasUserRightsOnGuest(c, userId, eventId, guestId)

	if err != nil || hasRightError != nil || hasRight == false {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	ticket, creationError := queries.CreateGuestWallet(ctx, db.CreateGuestWalletParams{
		GuestsID:      int32(body.GuestID),
		WalletsTypeID: int32(body.WalletTypeID),
		Token:         uuid.NewString(),
		Balance: sql.NullInt32{
			Int32: 0,
			Valid: true,
		},
	})

	if creationError != nil {
		return c.Status(fiber.StatusBadRequest).JSON(creationError.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(ticket)
}

func DeleteGuestWallet(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)

	eventId, err := strconv.Atoi(c.AllParams()["eventId"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	walletId, err := strconv.Atoi(c.AllParams()["walletId"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	guestId, err := strconv.Atoi(c.AllParams()["guestId"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userId, err := jwtauth.GetCurrentUserId(c)
	hasRight, hasRightError := utils.HasUserRightsOnGuest(c, userId, eventId, guestId)

	if err != nil || hasRightError != nil || hasRight == false {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	deletionError := queries.DeleteGuestWallet(ctx, int32(walletId))

	if deletionError != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusCreated)
}
