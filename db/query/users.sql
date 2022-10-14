-- name: CreateUser :one
INSERT INTO users (
  username, email, password, firstname, lastname, phone
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1;

