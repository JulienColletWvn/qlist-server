-- name: CreateCashier :one
INSERT INTO cashiers (
  users_id, events_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteCashier :exec
DELETE FROM cashiers
WHERE id = $1;



