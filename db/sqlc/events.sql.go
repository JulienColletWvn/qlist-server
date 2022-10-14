// Code generated by sqlc. DO NOT EDIT.
// source: events.sql

package db

import (
	"context"
	"time"
)

const createEvent = `-- name: CreateEvent :one
INSERT INTO events (
  name, description, start_date, end_date, location, free_wifi, public, tickets_amount
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING id, created_at, updated_at, name, description, start_date, end_date, location, free_wifi, public, tickets_amount
`

type CreateEventParams struct {
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	Location      string    `json:"location"`
	FreeWifi      bool      `json:"free_wifi"`
	Public        bool      `json:"public"`
	TicketsAmount int32     `json:"tickets_amount"`
}

func (q *Queries) CreateEvent(ctx context.Context, arg CreateEventParams) (Event, error) {
	row := q.db.QueryRowContext(ctx, createEvent,
		arg.Name,
		arg.Description,
		arg.StartDate,
		arg.EndDate,
		arg.Location,
		arg.FreeWifi,
		arg.Public,
		arg.TicketsAmount,
	)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Description,
		&i.StartDate,
		&i.EndDate,
		&i.Location,
		&i.FreeWifi,
		&i.Public,
		&i.TicketsAmount,
	)
	return i, err
}

const createEventAdministrator = `-- name: CreateEventAdministrator :one
INSERT INTO events_administrators (
  users_id, events_id
) VALUES (
  $1, $2
)
RETURNING id, created_at, users_id, events_id
`

type CreateEventAdministratorParams struct {
	UsersID  int32 `json:"users_id"`
	EventsID int32 `json:"events_id"`
}

func (q *Queries) CreateEventAdministrator(ctx context.Context, arg CreateEventAdministratorParams) (EventsAdministrator, error) {
	row := q.db.QueryRowContext(ctx, createEventAdministrator, arg.UsersID, arg.EventsID)
	var i EventsAdministrator
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UsersID,
		&i.EventsID,
	)
	return i, err
}

const createEventGuest = `-- name: CreateEventGuest :one
INSERT INTO events_guests (
  events_id, guests_id
) VALUES (
  $1, $2
)
RETURNING id, guests_id, events_id
`

type CreateEventGuestParams struct {
	EventsID int32 `json:"events_id"`
	GuestsID int32 `json:"guests_id"`
}

func (q *Queries) CreateEventGuest(ctx context.Context, arg CreateEventGuestParams) (EventsGuest, error) {
	row := q.db.QueryRowContext(ctx, createEventGuest, arg.EventsID, arg.GuestsID)
	var i EventsGuest
	err := row.Scan(&i.ID, &i.GuestsID, &i.EventsID)
	return i, err
}

const createEventGuestGroup = `-- name: CreateEventGuestGroup :one
INSERT INTO events_guests_groups (
  events_id, guests_groups_types
) VALUES (
  $1, $2
)
RETURNING id, guests_groups_types, events_id
`

type CreateEventGuestGroupParams struct {
	EventsID          int32 `json:"events_id"`
	GuestsGroupsTypes int32 `json:"guests_groups_types"`
}

func (q *Queries) CreateEventGuestGroup(ctx context.Context, arg CreateEventGuestGroupParams) (EventsGuestsGroup, error) {
	row := q.db.QueryRowContext(ctx, createEventGuestGroup, arg.EventsID, arg.GuestsGroupsTypes)
	var i EventsGuestsGroup
	err := row.Scan(&i.ID, &i.GuestsGroupsTypes, &i.EventsID)
	return i, err
}

const createEventPhoto = `-- name: CreateEventPhoto :one
INSERT INTO events_photos (
  events_id, url
) VALUES (
  $1, $2
)
RETURNING id, created_at, updated_at, events_id, url
`

type CreateEventPhotoParams struct {
	EventsID int32  `json:"events_id"`
	Url      string `json:"url"`
}

func (q *Queries) CreateEventPhoto(ctx context.Context, arg CreateEventPhotoParams) (EventsPhoto, error) {
	row := q.db.QueryRowContext(ctx, createEventPhoto, arg.EventsID, arg.Url)
	var i EventsPhoto
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.EventsID,
		&i.Url,
	)
	return i, err
}

const deleteEvent = `-- name: DeleteEvent :exec
DELETE FROM events
WHERE id = $1
`

func (q *Queries) DeleteEvent(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteEvent, id)
	return err
}

const deleteEventAdministrator = `-- name: DeleteEventAdministrator :exec
DELETE FROM events_administrators
WHERE users_id = $1
`

func (q *Queries) DeleteEventAdministrator(ctx context.Context, usersID int32) error {
	_, err := q.db.ExecContext(ctx, deleteEventAdministrator, usersID)
	return err
}

const deleteEventGuest = `-- name: DeleteEventGuest :exec
DELETE FROM events_guests
WHERE id = $1
`

func (q *Queries) DeleteEventGuest(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteEventGuest, id)
	return err
}

const deleteEventGuestGroup = `-- name: DeleteEventGuestGroup :exec
DELETE FROM events_guests_groups
WHERE id = $1
`

func (q *Queries) DeleteEventGuestGroup(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteEventGuestGroup, id)
	return err
}

const deleteEventPhoto = `-- name: DeleteEventPhoto :exec
DELETE FROM events_photos
WHERE id = $1
`

func (q *Queries) DeleteEventPhoto(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteEventPhoto, id)
	return err
}

const getPublicEvents = `-- name: GetPublicEvents :many
SELECT id, created_at, updated_at, name, description, start_date, end_date, location, free_wifi, public, tickets_amount FROM events
WHERE public = true
ORDER BY start_date DESC
LIMIT $1
OFFSET $2
`

type GetPublicEventsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetPublicEvents(ctx context.Context, arg GetPublicEventsParams) ([]Event, error) {
	rows, err := q.db.QueryContext(ctx, getPublicEvents, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Event
	for rows.Next() {
		var i Event
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Description,
			&i.StartDate,
			&i.EndDate,
			&i.Location,
			&i.FreeWifi,
			&i.Public,
			&i.TicketsAmount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserEvents = `-- name: GetUserEvents :many
SELECT id, created_at, updated_at, name, description, start_date, end_date, location, free_wifi, public, tickets_amount FROM events
WHERE creator_id = $1
ORDER BY start_date DESC
LIMIT $1
OFFSET $2
`

type GetUserEventsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetUserEvents(ctx context.Context, arg GetUserEventsParams) ([]Event, error) {
	rows, err := q.db.QueryContext(ctx, getUserEvents, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Event
	for rows.Next() {
		var i Event
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.Description,
			&i.StartDate,
			&i.EndDate,
			&i.Location,
			&i.FreeWifi,
			&i.Public,
			&i.TicketsAmount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateEventDates = `-- name: UpdateEventDates :one
UPDATE events
SET start_date = $2, end_date = $3, updated_at = now()
WHERE id = $1
RETURNING id, created_at, updated_at, name, description, start_date, end_date, location, free_wifi, public, tickets_amount
`

type UpdateEventDatesParams struct {
	ID        int32     `json:"id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

func (q *Queries) UpdateEventDates(ctx context.Context, arg UpdateEventDatesParams) (Event, error) {
	row := q.db.QueryRowContext(ctx, updateEventDates, arg.ID, arg.StartDate, arg.EndDate)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Description,
		&i.StartDate,
		&i.EndDate,
		&i.Location,
		&i.FreeWifi,
		&i.Public,
		&i.TicketsAmount,
	)
	return i, err
}

const updateEventDescription = `-- name: UpdateEventDescription :one
UPDATE events
SET description = $2, updated_at = now()
WHERE id = $1
RETURNING id, created_at, updated_at, name, description, start_date, end_date, location, free_wifi, public, tickets_amount
`

type UpdateEventDescriptionParams struct {
	ID          int32  `json:"id"`
	Description string `json:"description"`
}

func (q *Queries) UpdateEventDescription(ctx context.Context, arg UpdateEventDescriptionParams) (Event, error) {
	row := q.db.QueryRowContext(ctx, updateEventDescription, arg.ID, arg.Description)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Description,
		&i.StartDate,
		&i.EndDate,
		&i.Location,
		&i.FreeWifi,
		&i.Public,
		&i.TicketsAmount,
	)
	return i, err
}

const updateEventIsPublic = `-- name: UpdateEventIsPublic :one
UPDATE events
SET public = $2, updated_at = now()
WHERE id = $1
RETURNING id, created_at, updated_at, name, description, start_date, end_date, location, free_wifi, public, tickets_amount
`

type UpdateEventIsPublicParams struct {
	ID     int32 `json:"id"`
	Public bool  `json:"public"`
}

func (q *Queries) UpdateEventIsPublic(ctx context.Context, arg UpdateEventIsPublicParams) (Event, error) {
	row := q.db.QueryRowContext(ctx, updateEventIsPublic, arg.ID, arg.Public)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Description,
		&i.StartDate,
		&i.EndDate,
		&i.Location,
		&i.FreeWifi,
		&i.Public,
		&i.TicketsAmount,
	)
	return i, err
}

const updateEventLocation = `-- name: UpdateEventLocation :one
UPDATE events
SET location = $2, updated_at = now()
WHERE id = $1
RETURNING id, created_at, updated_at, name, description, start_date, end_date, location, free_wifi, public, tickets_amount
`

type UpdateEventLocationParams struct {
	ID       int32  `json:"id"`
	Location string `json:"location"`
}

func (q *Queries) UpdateEventLocation(ctx context.Context, arg UpdateEventLocationParams) (Event, error) {
	row := q.db.QueryRowContext(ctx, updateEventLocation, arg.ID, arg.Location)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Description,
		&i.StartDate,
		&i.EndDate,
		&i.Location,
		&i.FreeWifi,
		&i.Public,
		&i.TicketsAmount,
	)
	return i, err
}

const updateEventName = `-- name: UpdateEventName :one
UPDATE events
SET name = $2, updated_at = now()
WHERE id = $1
RETURNING id, created_at, updated_at, name, description, start_date, end_date, location, free_wifi, public, tickets_amount
`

type UpdateEventNameParams struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdateEventName(ctx context.Context, arg UpdateEventNameParams) (Event, error) {
	row := q.db.QueryRowContext(ctx, updateEventName, arg.ID, arg.Name)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Description,
		&i.StartDate,
		&i.EndDate,
		&i.Location,
		&i.FreeWifi,
		&i.Public,
		&i.TicketsAmount,
	)
	return i, err
}

const updateEventTicketsAmount = `-- name: UpdateEventTicketsAmount :one
UPDATE events
SET tickets_amount = $2, updated_at = now()
WHERE id = $1
RETURNING id, created_at, updated_at, name, description, start_date, end_date, location, free_wifi, public, tickets_amount
`

type UpdateEventTicketsAmountParams struct {
	ID            int32 `json:"id"`
	TicketsAmount int32 `json:"tickets_amount"`
}

func (q *Queries) UpdateEventTicketsAmount(ctx context.Context, arg UpdateEventTicketsAmountParams) (Event, error) {
	row := q.db.QueryRowContext(ctx, updateEventTicketsAmount, arg.ID, arg.TicketsAmount)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Description,
		&i.StartDate,
		&i.EndDate,
		&i.Location,
		&i.FreeWifi,
		&i.Public,
		&i.TicketsAmount,
	)
	return i, err
}

const updateEventWifiAvailability = `-- name: UpdateEventWifiAvailability :one
UPDATE events
SET free_wifi = $2, updated_at = now()
WHERE id = $1
RETURNING id, created_at, updated_at, name, description, start_date, end_date, location, free_wifi, public, tickets_amount
`

type UpdateEventWifiAvailabilityParams struct {
	ID       int32 `json:"id"`
	FreeWifi bool  `json:"free_wifi"`
}

func (q *Queries) UpdateEventWifiAvailability(ctx context.Context, arg UpdateEventWifiAvailabilityParams) (Event, error) {
	row := q.db.QueryRowContext(ctx, updateEventWifiAvailability, arg.ID, arg.FreeWifi)
	var i Event
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Description,
		&i.StartDate,
		&i.EndDate,
		&i.Location,
		&i.FreeWifi,
		&i.Public,
		&i.TicketsAmount,
	)
	return i, err
}