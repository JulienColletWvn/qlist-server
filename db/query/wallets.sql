-- name: CreateGuestWallet :one
INSERT INTO wallets (
        guests_id,
        wallets_type_id,
        token,
        balance
    )
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: DeleteGuestWallet :exec
DELETE FROM wallets
WHERE id = $1;
-- name: UpdateGuestWalletBalance :one
UPDATE wallets
SET balance = $1
WHERE id = $2
RETURNING *;
-- name: UpdateGuestWalletToken :one
UPDATE wallets
SET token = $1
WHERE id = $2
RETURNING *;