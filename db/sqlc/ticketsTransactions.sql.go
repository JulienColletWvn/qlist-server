// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: ticketsTransactions.sql

package db

import (
	"context"
)

const createTicketTransaction = `-- name: CreateTicketTransaction :one
INSERT INTO tickets_transactions (
        tickets_id,
        stewards_id,
        amount,
        status
    )
VALUES ($1, $2, $3, $4)
RETURNING id, created_at, tickets_id, stewards_id, amount, status
`

type CreateTicketTransactionParams struct {
	TicketsID  int32             `json:"tickets_id"`
	StewardsID int32             `json:"stewards_id"`
	Amount     int32             `json:"amount"`
	Status     TransactionStatus `json:"status"`
}

func (q *Queries) CreateTicketTransaction(ctx context.Context, arg CreateTicketTransactionParams) (TicketsTransaction, error) {
	row := q.db.QueryRow(ctx, createTicketTransaction,
		arg.TicketsID,
		arg.StewardsID,
		arg.Amount,
		arg.Status,
	)
	var i TicketsTransaction
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.TicketsID,
		&i.StewardsID,
		&i.Amount,
		&i.Status,
	)
	return i, err
}

const updateTicketTransactionStatus = `-- name: UpdateTicketTransactionStatus :one
UPDATE tickets_transactions
SET status = $1
WHERE tickets_transactions.id = $2
RETURNING id, created_at, tickets_id, stewards_id, amount, status
`

type UpdateTicketTransactionStatusParams struct {
	Status TransactionStatus `json:"status"`
	ID     int32             `json:"id"`
}

func (q *Queries) UpdateTicketTransactionStatus(ctx context.Context, arg UpdateTicketTransactionStatusParams) (TicketsTransaction, error) {
	row := q.db.QueryRow(ctx, updateTicketTransactionStatus, arg.Status, arg.ID)
	var i TicketsTransaction
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.TicketsID,
		&i.StewardsID,
		&i.Amount,
		&i.Status,
	)
	return i, err
}