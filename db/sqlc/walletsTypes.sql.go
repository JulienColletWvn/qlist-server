// Code generated by sqlc. DO NOT EDIT.
// source: walletsTypes.sql

package db

import (
	"context"
	"time"
)

const createWalletType = `-- name: CreateWalletType :one
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

type CreateWalletTypeParams struct {
	EventsID          int32     `json:"events_id"`
	Name              string    `json:"name"`
	StartValidityDate time.Time `json:"start_validity_date"`
	EndValidityDate   time.Time `json:"end_validity_date"`
	MaxAmount         int32     `json:"max_amount"`
	OnlineReload      bool      `json:"online_reload"`
}

func (q *Queries) CreateWalletType(ctx context.Context, arg CreateWalletTypeParams) (WalletsType, error) {
	row := q.db.QueryRowContext(ctx, createWalletType,
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

const deleteWalletType = `-- name: DeleteWalletType :exec
DELETE FROM wallets_types
WHERE wallets_types.id = $1
    AND wallets_types.events_id IN (
        SELECT events_administrators.events.id
        FROM events_administrators
        WHERE events_administrators.users_id = $2
    )
`

type DeleteWalletTypeParams struct {
	ID      int32 `json:"id"`
	UsersID int32 `json:"users_id"`
}

func (q *Queries) DeleteWalletType(ctx context.Context, arg DeleteWalletTypeParams) error {
	_, err := q.db.ExecContext(ctx, deleteWalletType, arg.ID, arg.UsersID)
	return err
}