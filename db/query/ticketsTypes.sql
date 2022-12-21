-- name: CreateEventTicketType :one
INSERT INTO tickets_types (
        events_id,
        name,
        start_validity_date,
        end_validity_date,
        usage_limitation,
        usage_unlimited
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
-- name: GetEventTicketsTypes :many
SELECT *
FROM tickets_types
WHERE tickets_types.events_id IN (
        SELECT events_administrators.events_id
        FROM events_administrators
        WHERE events_administrators.events_id = $1
            AND events_administrators.users_id = $2
    );
-- name: GetEventTicketsType :one
SELECT *
FROM tickets_types
WHERE tickets_types.id = $1
    AND tickets_types.events_id IN (
        SELECT events_administrators.events_id
        FROM events_administrators
        WHERE events_administrators.users_id = $2
    );
-- name: DeleteEventTicketsType :exec
DELETE FROM tickets_types
WHERE tickets_types.id = $1
    AND tickets_types.events_id IN (
        SELECT events_administrators.events_id
        FROM events_administrators
        WHERE events_administrators.users_id = $2
    );