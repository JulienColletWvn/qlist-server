-- name: CreateEventGuest :one
INSERT INTO guests (note, events_id, contacts_id)
VALUES ($1, $2, $3)
RETURNING *;
-- name: CreateEventGuests :copyfrom
INSERT INTO guests (note, events_id, contacts_id)
VALUES ($1, $2, $3);
-- name: DeleteUserEventGuest :exec
DELETE FROM guests
WHERE guests.id = $1
    AND guests.events_id IN (
        SELECT events_administrators.events.id
        FROM events_administrators
        WHERE events_administrators.users_id = $2
    );
-- name: GetUserEventGuest :one
SELECT *
FROM guests
WHERE guests.id = $1
    AND guests.events_id IN (
        SELECT events_administrators.events.id
        FROM events_administrators
        WHERE events_administrators.events_id = $2
            AND events_administrators.users_id = $3
    );
-- name: GetUserEventGuests :one
SELECT *
FROM guests
WHERE guests.events_id IN (
        SELECT events_administrators.events.id
        FROM events_administrators
        WHERE events_administrators.events_id = $1
            AND events_administrators.users_id = $2
    );