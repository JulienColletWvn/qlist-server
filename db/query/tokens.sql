-- name: CreateToken :one
INSERT INTO tokens (
  uuid, wallets_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteToken :exec
DELETE FROM tokens
WHERE id = $1;

-- name: CreateTokenTransaction :one
INSERT INTO tokens_transactions (
  transaction_date, amount, online_sell, cashiers_id, sellers_id, tokens_id, events_products_id, status
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: DeleteTokenTransaction :exec
DELETE FROM tokens_transactions
WHERE id = $1;

