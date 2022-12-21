-- name: CreateEventProduct :one
INSERT INTO events_products (events_id, name, tokens_amount_pricing)
VALUES ($1, $2, $3)
RETURNING *;
-- name: CreateEventProductSeller :one
INSERT INTO events_products_sellers (sellers_id, events_products_id)
VALUES ($1, $2)
RETURNING *;
-- name: DeleteEventProduct :exec
DELETE FROM events_products
WHERE id = $1;
-- name: DeleteEventProductSeller :exec
DELETE FROM events_products_sellers
WHERE id = $1;