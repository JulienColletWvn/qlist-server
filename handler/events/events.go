package handler

import (
	"context"
	"fmt"
	db "qlist/db/sqlc"
	jwtauth "qlist/middleware"
	"qlist/utils"
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

func GetUserEvents(c *fiber.Ctx) error {
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
				Type:    string(content.Type),
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
	ctx := context.Background()
	userId, err := jwtauth.GetCurrentUserId(c)
	params := CreateEventParams{}
	queries := db.New(utils.Database)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(err)
	}

	if err := c.BodyParser(&params); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	if err := utils.ValidateStruct(params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	event, err := queries.CreateEvent(ctx, db.CreateEventParams{
		StartDate:     params.StartDate,
		EndDate:       params.EndDate,
		Location:      params.Location,
		FreeWifi:      params.FreeWifi,
		Public:        params.Public,
		TicketsAmount: params.TicketsAmount,
		CreatorID:     int32(userId),
	})

	fmt.Println(event)
	fmt.Println(err)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	for _, content := range params.EventContents {
		_, err := queries.CreateEventContent(ctx, db.CreateEventContentParams{
			Type:     db.EventsContentType(content.Type),
			Content:  content.Content,
			Lang:     content.Lang,
			EventsID: event.ID,
		})

		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
	}

	_, adminCreationError := queries.CreateEventAdministrator(ctx, db.CreateEventAdministratorParams{
		EventsID: event.ID,
		UsersID:  int32(userId),
	})

	if adminCreationError != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

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
