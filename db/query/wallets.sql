-- name: CreateWallet :one
INSERT INTO wallets (
  guests_id, wallets_type_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteWallet :exec
DELETE FROM wallets
WHERE id = $1;

-- name: CreateWalletType :one
INSERT INTO wallets_type (
  events_id, name, start_validity_date, end_validity_date, max_amount, online_reload
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: DeleteWalletType :exec
DELETE FROM wallets_type
WHERE id = $1;

-- name: CreateWalletPricing :one
INSERT INTO wallets_pricing (
  quantity, unit_price, wallets_type_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteWalletPricing :exec
DELETE FROM wallets_pricing
WHERE id = $1;

-- name: CreateWalletTransaction :one
INSERT INTO wallets_transactions (
  cashiers_id, wallets_id, wallets_pricing_id, units_sold, status
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: DeleteWalletTransaction :exec
DELETE FROM wallets_transactions
WHERE id = $1;




