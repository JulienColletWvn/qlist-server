-- name: GetGuestTicketTransactions :many
SELECT *
FROM tickets_transactions
WHERE tickets_id = $1;
-- name: GetGuestTicketTransaction :one
SELECT *
FROM tickets_transactions
WHERE tickets_id = $1
    and id = $2;
-- name: CreateGuestTicketTransaction :one
INSERT INTO tickets_transactions (
        tickets_id,
        stewards_id,
        amount,
        status
    )
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: UpdateGuestTicketTransactionStatus :one
UPDATE tickets_transactions
SET status = $1
WHERE tickets_transactions.id = $2
RETURNING *;