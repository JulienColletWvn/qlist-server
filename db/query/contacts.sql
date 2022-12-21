-- name: CreateUserContacts :copyfrom
INSERT INTO contacts (
    email,
    firstname,
    lastname,
    phone,
    lang,
    creator_id
  )
VALUES ($1, $2, $3, $4, $5, $6);
-- name: GetUserContacts :many
SELECT *
FROM contacts
WHERE creator_id = $1;
-- name: GetUserContact :many
SELECT *
FROM contacts
WHERE creator_id = $1
  AND id = $2;
-- name: DeleteUserContact :exec
DELETE FROM contacts
WHERE creator_id = $1
  AND id = $2;