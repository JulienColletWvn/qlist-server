-- name: CreateContact :one
INSERT INTO contacts (
    email,
    firstname,
    lastname,
    phone,
    lang,
    creator_id
  )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
-- name: GetContacts :one
SELECT *
FROM contacts
WHERE creator_id = $1;
-- name: GetContact :many
SELECT *
FROM contacts
WHERE creator_id = $1
  AND id = $2;
-- name: DeleteContact :exec
DELETE FROM contacts
WHERE creator_id = $1
  AND id = $2;