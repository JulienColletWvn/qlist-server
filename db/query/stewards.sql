-- name: CreateUserEventSteward :one
INSERT INTO stewards (events_id, contacts_id)
VALUES ($1, $2)
RETURNING *;
-- name: DeleteUserEventSteward :exec
DELETE FROM stewards
WHERE stewards.id = $1
    AND stewards.events_id IN (
        SELECT events_administrators.events.id
        FROM events_administrators
        WHERE events_administrators.users_id = $2
    );
-- name: GetUserEventSteward :one
SELECT *
FROM stewards
WHERE stewards.id = $1
    AND stewards.events_id IN (
        SELECT events_administrators.events.id
        FROM events_administrators
        WHERE events_administrators.events_id = $2
            AND events_administrators.users_id = $3
    );
-- name: GetUserEventStewards :one
SELECT *
FROM stewards
WHERE stewards.events_id IN (
        SELECT events_administrators.events.id
        FROM events_administrators
        WHERE events_administrators.events_id = $1
            AND events_administrators.users_id = $2
    );