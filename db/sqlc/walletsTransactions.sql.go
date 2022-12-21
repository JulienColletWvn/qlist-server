// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: walletsTransactions.sql

package db

import (
	"context"
	"database/sql"
)

const createWalletTransaction = `-- name: CreateWalletTransaction :one
INSERT INTO wallets_transactions (
        amount,
        online_sell,
        cashiers_id,
        sellers_id,
        events_products_id,
        status
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, created_at, updated_at, amount, online_sell, cashiers_id, sellers_id, events_products_id, status
`

type CreateWalletTransactionParams struct {
	Amount           int32             `json:"amount"`
	OnlineSell       bool              `json:"online_sell"`
	CashiersID       sql.NullInt32     `json:"cashiers_id"`
	SellersID        sql.NullInt32     `json:"sellers_id"`
	EventsProductsID int32             `json:"events_products_id"`
	Status           TransactionStatus `json:"status"`
}

func (q *Queries) CreateWalletTransaction(ctx context.Context, arg CreateWalletTransactionParams) (WalletsTransaction, error) {
	row := q.db.QueryRow(ctx, createWalletTransaction,
		arg.Amount,
		arg.OnlineSell,
		arg.CashiersID,
		arg.SellersID,
		arg.EventsProductsID,
		arg.Status,
	)
	var i WalletsTransaction
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Amount,
		&i.OnlineSell,
		&i.CashiersID,
		&i.SellersID,
		&i.EventsProductsID,
		&i.Status,
	)
	return i, err
}

const updateWalletTransactionStatus = `-- name: UpdateWalletTransactionStatus :one
UPDATE wallets_transactions
SET status = $1
WHERE wallets_transactions.id = $2
RETURNING id, created_at, updated_at, amount, online_sell, cashiers_id, sellers_id, events_products_id, status
`

type UpdateWalletTransactionStatusParams struct {
	Status TransactionStatus `json:"status"`
	ID     int32             `json:"id"`
}

func (q *Queries) UpdateWalletTransactionStatus(ctx context.Context, arg UpdateWalletTransactionStatusParams) (WalletsTransaction, error) {
	row := q.db.QueryRow(ctx, updateWalletTransactionStatus, arg.Status, arg.ID)
	var i WalletsTransaction
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Amount,
		&i.OnlineSell,
		&i.CashiersID,
		&i.SellersID,
		&i.EventsProductsID,
		&i.Status,
	)
	return i, err
}
