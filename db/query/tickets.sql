-- name: CreateGuestTicket :one
INSERT INTO tickets (tickets_type_id, guests_id)
VALUES ($1, $2)
RETURNING *;
-- name: GetGuestTickets :many
SELECT *
FROM tickets
WHERE tickets.guests_id = $1;
-- name: GetGuestTicket :one
SELECT *
FROM tickets
WHERE tickets.guests_id = $1
    AND tickets.id = $2;
-- name: DeleteGuestTicket :exec
DELETE FROM tickets
WHERE tickets.id = $1;