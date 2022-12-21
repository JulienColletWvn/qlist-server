-- name: CreateTicketTransaction :one
INSERT INTO tickets_transactions (
        tickets_id,
        stewards_id,
        amount,
        status
    )
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: UpdateTicketTransactionStatus :one
UPDATE tickets_transactions
SET status = $1
WHERE tickets_transactions.id = $2
RETURNING *;