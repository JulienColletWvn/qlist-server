// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: cashiers.sql

package db

import (
	"context"
)

const createUserEventCashier = `-- name: CreateUserEventCashier :one
INSERT INTO cashiers (events_id, contacts_id)
VALUES ($1, $2)
RETURNING id, created_at, contacts_id, events_id
`

type CreateUserEventCashierParams struct {
	EventsID   int32 `json:"events_id"`
	ContactsID int32 `json:"contacts_id"`
}

func (q *Queries) CreateUserEventCashier(ctx context.Context, arg CreateUserEventCashierParams) (Cashier, error) {
	row := q.db.QueryRow(ctx, createUserEventCashier, arg.EventsID, arg.ContactsID)
	var i Cashier
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.ContactsID,
		&i.EventsID,
	)
	return i, err
}

const deleteUserEventCashier = `-- name: DeleteUserEventCashier :exec
DELETE FROM cashiers
WHERE cashiers.id = $1
    AND cashiers.events_id IN (
        SELECT events_administrators.events.id
        FROM events_administrators
        WHERE events_administrators.users_id = $2
    )
`

type DeleteUserEventCashierParams struct {
	ID      int32 `json:"id"`
	UsersID int32 `json:"users_id"`
}

func (q *Queries) DeleteUserEventCashier(ctx context.Context, arg DeleteUserEventCashierParams) error {
	_, err := q.db.Exec(ctx, deleteUserEventCashier, arg.ID, arg.UsersID)
	return err
}

const getUserEventCashier = `-- name: GetUserEventCashier :one
SELECT id, created_at, contacts_id, events_id
FROM cashiers
WHERE cashiers.id = $1
    AND cashiers.events_id IN (
        SELECT events_administrators.events.id
        FROM events_administrators
        WHERE events_administrators.events_id = $2
            AND events_administrators.users_id = $3
    )
`

type GetUserEventCashierParams struct {
	ID       int32 `json:"id"`
	EventsID int32 `json:"events_id"`
	UsersID  int32 `json:"users_id"`
}

func (q *Queries) GetUserEventCashier(ctx context.Context, arg GetUserEventCashierParams) (Cashier, error) {
	row := q.db.QueryRow(ctx, getUserEventCashier, arg.ID, arg.EventsID, arg.UsersID)
	var i Cashier
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.ContactsID,
		&i.EventsID,
	)
	return i, err
}

const getUserEventCashiers = `-- name: GetUserEventCashiers :one
SELECT id, created_at, contacts_id, events_id
FROM cashiers
WHERE cashiers.events_id IN (
        SELECT events_administrators.events.id
        FROM events_administrators
        WHERE events_administrators.events_id = $1
            AND events_administrators.users_id = $2
    )
`

type GetUserEventCashiersParams struct {
	EventsID int32 `json:"events_id"`
	UsersID  int32 `json:"users_id"`
}

func (q *Queries) GetUserEventCashiers(ctx context.Context, arg GetUserEventCashiersParams) (Cashier, error) {
	row := q.db.QueryRow(ctx, getUserEventCashiers, arg.EventsID, arg.UsersID)
	var i Cashier
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.ContactsID,
		&i.EventsID,
	)
	return i, err
}
