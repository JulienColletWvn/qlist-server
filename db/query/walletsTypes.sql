-- name: GetEventWalletsTypes :many
SELECT *
FROM wallets_types
WHERE wallets_types.events_id IN (
        SELECT events_administrators.events_id
        FROM events_administrators
        WHERE events_administrators.events_id = $1
            AND events_administrators.users_id = $2
    );
-- name: GetEventWalletsType :one
SELECT *
FROM wallets_types
WHERE wallets_types.id = $1
    AND wallets_types.events_id IN (
        SELECT events_administrators.events_id
        FROM events_administrators
        WHERE events_administrators.users_id = $2
    );
-- name: CreateEventWalletType :one
INSERT INTO wallets_types (
        events_id,
        name,
        start_validity_date,
        end_validity_date,
        max_amount,
        online_reload
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
-- name: DeleteEventWalletType :exec
DELETE FROM wallets_types
WHERE wallets_types.id = $1
    AND wallets_types.events_id IN (
        SELECT events_administrators.events.id
        FROM events_administrators
        WHERE events_administrators.users_id = $2
    );