-- name: CreateUser :one
INSERT INTO users (
    username,
    email,
    password,
    firstname,
    lastname,
    phone
  )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
-- name: GetUserById :one
SELECT *
FROM users
WHERE id = $1;
-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;