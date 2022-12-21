-- name: CreateWalletTransaction :one
INSERT INTO wallets_transactions (
        amount,
        online_sell,
        cashiers_id,
        sellers_id,
        events_products_id,
        status
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
-- name: UpdateWalletTransactionStatus :one
UPDATE wallets_transactions
SET status = $1
WHERE wallets_transactions.id = $2
RETURNING *;