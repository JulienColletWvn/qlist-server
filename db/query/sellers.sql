-- name: CreateSeller :one
INSERT INTO sellers (events_id, contacts_id)
VALUES ($1, $2)
RETURNING *;
-- name: DeleteSeller :exec
DELETE FROM sellers
WHERE sellers.id = $1
    AND sellers.events_id IN (
        SELECT events_administrators.events.id
        FROM events_administrators
        WHERE events_administrators.users_id = $2
    );
-- name: GetSeller :one
SELECT *
FROM sellers
WHERE sellers.id = $1
    AND sellers.events_id IN (
        SELECT events_administrators.events.id
        FROM events_administrators
        WHERE events_administrators.events_id = $2
            AND events_administrators.users_id = $3
    );
-- name: GetSellers :one
SELECT *
FROM sellers
WHERE sellers.events_id IN (
        SELECT events_administrators.events.id
        FROM events_administrators
        WHERE events_administrators.events_id = $1
            AND events_administrators.users_id = $2
    );