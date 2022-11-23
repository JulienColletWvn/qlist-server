-- name: CreateWalletPricing :one
INSERT INTO wallets_pricings (
        quantity,
        unit_price,
        wallets_type_id
    )
VALUES ($1, $2, $3)
RETURNING *;
-- name: DeleteWalletPricing :exec
DELETE FROM wallets_pricings p USING wallets_types t
WHERE p.id = $1
    AND p.wallets_type_id = t.id
    AND t.events_id IN (
        SELECT events_administrators.events.id
        FROM events_administrators
        WHERE events_administrators.users_id = $2
    );