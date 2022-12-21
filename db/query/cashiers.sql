-- name: CreateUserEventCashier :one
INSERT INTO cashiers (events_id, contacts_id)
VALUES ($1, $2)
RETURNING *;
-- name: DeleteUserEventCashier :exec
DELETE FROM cashiers
WHERE cashiers.id = $1
    AND cashiers.events_id IN (
        SELECT events_administrators.events.id
        FROM events_administrators
        WHERE events_administrators.users_id = $2
    );
-- name: GetUserEventCashier :one
SELECT *
FROM cashiers
WHERE cashiers.id = $1
    AND cashiers.events_id IN (
        SELECT events_administrators.events.id
        FROM events_administrators
        WHERE events_administrators.events_id = $2
            AND events_administrators.users_id = $3
    );
-- name: GetUserEventCashiers :one
SELECT *
FROM cashiers
WHERE cashiers.events_id IN (
        SELECT events_administrators.events.id
        FROM events_administrators
        WHERE events_administrators.events_id = $1
            AND events_administrators.users_id = $2
    );