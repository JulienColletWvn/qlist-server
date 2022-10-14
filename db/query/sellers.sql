
-- name: CreateSeller :one
INSERT INTO sellers (
  users_id, events_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteSeller :exec
DELETE FROM sellers
WHERE id = $1;



