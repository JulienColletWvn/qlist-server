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

type EventContent struct {
	Type    string `json:"type"`
	Content string `json:"content"`
	Lang    string `json:"lang"`
}

type CreateEventParams struct {
	StartDate     time.Time      `json:"start_date"`
	EndDate       time.Time      `json:"end_date"`
	Location      string         `json:"location"`
	FreeWifi      bool           `json:"free_wifi"`
	Public        bool           `json:"public"`
	TicketsAmount int32          `json:"tickets_amount"`
	EventContents []EventContent `json:"contents"`
}

type EventResponse struct {
	Id            int            `json:"id"`
	StartDate     time.Time      `json:"start_date"`
	EndDate       time.Time      `json:"end_date"`
	Location      string         `json:"location"`
	FreeWifi      bool           `json:"free_wifi"`
	Public        bool           `json:"public"`
	TicketsAmount int32          `json:"tickets_amount"`
	Contents      []EventContent `json:"contents"`
	CreatedAt     string         `json:"created_at"`
	UpdatedAt     string         `json:"updated_at"`
}

func GetEvents(c *fiber.Ctx) error {
	ctx := context.Background()
	userId, err := jwtauth.GetCurrentUserId(c)
	queries := db.New(utils.Database)
	eventsResponse := []EventResponse{}

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	events, err := queries.GetAdministratorEvents(ctx, int32(userId))

	if err != nil {
		c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	for _, event := range events {
		parsedContents := []EventContent{}
		contents, err := queries.GetEventContents(ctx, event.ID)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		for _, content := range contents {
			parsedContents = append(parsedContents, EventContent{
				Content: content.Content,
				Lang:    content.Lang,
				Type:    string(content.Type.EventsContentType),
			})
		}

		eventsResponse = append(eventsResponse, EventResponse{
			Id:            int(event.ID),
			CreatedAt:     event.CreatedAt.Time.String(),
			UpdatedAt:     event.UpdatedAt.Time.String(),
			StartDate:     event.StartDate.UTC(),
			EndDate:       event.EndDate.UTC(),
			Location:      event.Location,
			FreeWifi:      event.FreeWifi,
			Public:        event.Public,
			TicketsAmount: event.TicketsAmount,
			Contents:      parsedContents,
		})
	}

	return c.Status(fiber.StatusOK).JSON(eventsResponse)
}

func CreateEvent(c *fiber.Ctx) error {
	params := CreateEventParams{}
	ctx := context.Background()

	userId, err := jwtauth.GetCurrentUserId(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	tx, err := utils.Database.Begin(ctx)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	defer tx.Rollback(ctx)

	qtx := db.New(utils.Database).WithTx(tx)

	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	if err := utils.ValidateStruct(params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	event, err := qtx.CreateEvent(ctx, db.CreateEventParams{
		StartDate:     params.StartDate,
		EndDate:       params.EndDate,
		Location:      params.Location,
		FreeWifi:      params.FreeWifi,
		Public:        params.Public,
		TicketsAmount: params.TicketsAmount,
		CreatorID:     int32(userId),
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	for _, content := range params.EventContents {
		_, err := qtx.CreateEventContent(ctx, db.CreateEventContentParams{
			Type: db.NullEventsContentType{
				EventsContentType: db.EventsContentType(content.Type),
				Valid:             true,
			},
			Content:  content.Content,
			Lang:     content.Lang,
			EventsID: event.ID,
		})

		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	}

	_, adminCreationError := qtx.CreateEventAdministrator(ctx, db.CreateEventAdministratorParams{
		EventsID: event.ID,
		UsersID:  int32(userId),
	})

	if adminCreationError != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	tx.Commit(ctx)

	return c.Status(fiber.StatusOK).JSON(EventResponse{
		Id:            int(event.ID),
		CreatedAt:     event.CreatedAt.Time.String(),
		UpdatedAt:     event.UpdatedAt.Time.String(),
		StartDate:     event.StartDate.UTC(),
		EndDate:       event.EndDate.UTC(),
		Location:      event.Location,
		FreeWifi:      event.FreeWifi,
		Public:        event.Public,
		TicketsAmount: event.TicketsAmount,
		Contents:      params.EventContents,
	})
}

func GetEvent(c *fiber.Ctx) error {
	ctx := context.Background()
	queries := db.New(utils.Database)
	eventContents := []EventContent{}

	eventId, err := strconv.Atoi(c.AllParams()["eventId"])

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userId, err := jwtauth.GetCurrentUserId(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	event, err := queries.GetAdministratorEvent(ctx, db.GetAdministratorEventParams{
		Column1:  int32(userId),
		EventsID: int32(eventId),
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	contents, err := queries.GetEventContents(ctx, int32(eventId))

	for _, content := range contents {
		eventContents = append(eventContents, EventContent{
			Type:    string(content.Type.EventsContentType),
			Content: content.Content,
			Lang:    content.Lang,
		})
	}

	return c.Status(fiber.StatusOK).JSON(EventResponse{
		Id:            int(event.ID),
		StartDate:     event.StartDate.UTC(),
		EndDate:       event.EndDate.UTC(),
		Location:      event.Location,
		FreeWifi:      event.FreeWifi,
		Public:        event.Public,
		TicketsAmount: event.TicketsAmount,
		Contents:      eventContents,
		CreatedAt:     event.CreatedAt.Time.UTC().String(),
		UpdatedAt:     event.UpdatedAt.Time.UTC().String(),
	})
}

func DeleteUserEvent(c *fiber.Ctx) error {
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

	deletionError := queries.DeleteAdministratorEvent(ctx, db.DeleteAdministratorEventParams{
		ID:      int32(eventId),
		UsersID: int32(userId),
	})

	if deletionError != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	return c.SendStatus(fiber.StatusOK)
}
