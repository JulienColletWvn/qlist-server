// Code generated by sqlc. DO NOT EDIT.
// source: cashiers.sql

package db

import (
	"context"
)

const createCashier = `-- name: CreateCashier :one
INSERT INTO cashiers (
  users_id, events_id
) VALUES (
  $1, $2
)
RETURNING id, created_at, users_id, events_id
`

type CreateCashierParams struct {
	UsersID  int32 `json:"users_id"`
	EventsID int32 `json:"events_id"`
}

func (q *Queries) CreateCashier(ctx context.Context, arg CreateCashierParams) (Cashier, error) {
	row := q.db.QueryRowContext(ctx, createCashier, arg.UsersID, arg.EventsID)
	var i Cashier
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UsersID,
		&i.EventsID,
	)
	return i, err
}

const deleteCashier = `-- name: DeleteCashier :exec
DELETE FROM cashiers
WHERE id = $1
`

func (q *Queries) DeleteCashier(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteCashier, id)
	return err
}
