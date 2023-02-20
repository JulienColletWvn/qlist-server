// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: walletsTypes.sql

package db

import (
	"context"
	"time"
)

const createEventWalletType = `-- name: CreateEventWalletType :one
INSERT INTO wallets_types (
        events_id,
        name,
        start_validity_date,
        end_validity_date,
        max_amount,
        online_reload
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, created_at, updated_at, events_id, name, start_validity_date, end_validity_date, max_amount, online_reload
`

type CreateEventWalletTypeParams struct {
	EventsID          int32     `json:"events_id"`
	Name              string    `json:"name"`
	StartValidityDate time.Time `json:"start_validity_date"`
	EndValidityDate   time.Time `json:"end_validity_date"`
	MaxAmount         int32     `json:"max_amount"`
	OnlineReload      bool      `json:"online_reload"`
}

func (q *Queries) CreateEventWalletType(ctx context.Context, arg CreateEventWalletTypeParams) (WalletsType, error) {
	row := q.db.QueryRow(ctx, createEventWalletType,
		arg.EventsID,
		arg.Name,
		arg.StartValidityDate,
		arg.EndValidityDate,
		arg.MaxAmount,
		arg.OnlineReload,
	)
	var i WalletsType
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.EventsID,
		&i.Name,
		&i.StartValidityDate,
		&i.EndValidityDate,
		&i.MaxAmount,
		&i.OnlineReload,
	)
	return i, err
}

const deleteEventWalletType = `-- name: DeleteEventWalletType :exec
DELETE FROM wallets_types
WHERE wallets_types.id = $1
    AND wallets_types.events_id IN (
        SELECT events_administrators.events.id
        FROM events_administrators
        WHERE events_administrators.users_id = $2
    )
`

type DeleteEventWalletTypeParams struct {
	ID      int32 `json:"id"`
	UsersID int32 `json:"users_id"`
}

func (q *Queries) DeleteEventWalletType(ctx context.Context, arg DeleteEventWalletTypeParams) error {
	_, err := q.db.Exec(ctx, deleteEventWalletType, arg.ID, arg.UsersID)
	return err
}

const getEventWalletsType = `-- name: GetEventWalletsType :one
SELECT id, created_at, updated_at, events_id, name, start_validity_date, end_validity_date, max_amount, online_reload
FROM wallets_types
WHERE wallets_types.id = $1
    AND wallets_types.events_id IN (
        SELECT events_administrators.events_id
        FROM events_administrators
        WHERE events_administrators.users_id = $2
    )
`

type GetEventWalletsTypeParams struct {
	ID      int32 `json:"id"`
	UsersID int32 `json:"users_id"`
}

func (q *Queries) GetEventWalletsType(ctx context.Context, arg GetEventWalletsTypeParams) (WalletsType, error) {
	row := q.db.QueryRow(ctx, getEventWalletsType, arg.ID, arg.UsersID)
	var i WalletsType
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.EventsID,
		&i.Name,
		&i.StartValidityDate,
		&i.EndValidityDate,
		&i.MaxAmount,
		&i.OnlineReload,
	)
	return i, err
}

const getEventWalletsTypes = `-- name: GetEventWalletsTypes :many
SELECT id, created_at, updated_at, events_id, name, start_validity_date, end_validity_date, max_amount, online_reload
FROM wallets_types
WHERE wallets_types.events_id IN (
        SELECT events_administrators.events_id
        FROM events_administrators
        WHERE events_administrators.events_id = $1
            AND events_administrators.users_id = $2
    )
`

type GetEventWalletsTypesParams struct {
	EventsID int32 `json:"events_id"`
	UsersID  int32 `json:"users_id"`
}

func (q *Queries) GetEventWalletsTypes(ctx context.Context, arg GetEventWalletsTypesParams) ([]WalletsType, error) {
	rows, err := q.db.Query(ctx, getEventWalletsTypes, arg.EventsID, arg.UsersID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []WalletsType
	for rows.Next() {
		var i WalletsType
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.EventsID,
			&i.Name,
			&i.StartValidityDate,
			&i.EndValidityDate,
			&i.MaxAmount,
			&i.OnlineReload,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
