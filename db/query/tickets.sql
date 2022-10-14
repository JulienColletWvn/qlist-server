-- name: CreateTicket :one
INSERT INTO tickets (
  tickets_type_id, sellers_id, guests_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteTicket :exec
DELETE FROM tickets
WHERE id = $1;

-- name: CreateTicketType :one
INSERT INTO tickets_type (
  events_id, name, start_validity_date, end_validity_date, usage_limitation, usage_unlimited
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: DeleteTicketType :exec
DELETE FROM tickets_type
WHERE id = $1;

-- name: CreateTicketTransaction :one
INSERT INTO tickets_transactions (
  tickets_id, amount, status
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteTicketTransaction :exec
DELETE FROM tickets_transactions
WHERE id = $1;
