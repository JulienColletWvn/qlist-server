// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: walletsPricings.sql

package db

import (
	"context"
)

const createWalletPricing = `-- name: CreateWalletPricing :one
INSERT INTO wallets_pricings (
        quantity,
        unit_price,
        wallets_type_id
    )
VALUES ($1, $2, $3)
RETURNING id, type, quantity, unit_price, wallets_type_id
`

type CreateWalletPricingParams struct {
	Quantity      int32 `json:"quantity"`
	UnitPrice     int32 `json:"unit_price"`
	WalletsTypeID int32 `json:"wallets_type_id"`
}

func (q *Queries) CreateWalletPricing(ctx context.Context, arg CreateWalletPricingParams) (WalletsPricing, error) {
	row := q.db.QueryRow(ctx, createWalletPricing, arg.Quantity, arg.UnitPrice, arg.WalletsTypeID)
	var i WalletsPricing
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Quantity,
		&i.UnitPrice,
		&i.WalletsTypeID,
	)
	return i, err
}

const deleteWalletPricing = `-- name: DeleteWalletPricing :exec
DELETE FROM wallets_pricings p USING wallets_types t
WHERE p.id = $1
    AND p.wallets_type_id = t.id
    AND t.events_id IN (
        SELECT events_administrators.events.id
        FROM events_administrators
        WHERE events_administrators.users_id = $2
    )
`

type DeleteWalletPricingParams struct {
	ID      int32 `json:"id"`
	UsersID int32 `json:"users_id"`
}

func (q *Queries) DeleteWalletPricing(ctx context.Context, arg DeleteWalletPricingParams) error {
	_, err := q.db.Exec(ctx, deleteWalletPricing, arg.ID, arg.UsersID)
	return err
}
