-- name: CreateGuest :one
INSERT INTO guests (
  creator_id, email, firstname, lastname, phone
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: DeleteGuest :exec
DELETE FROM guests
WHERE id = $1;

-- name: CreateGuestGroup :one
INSERT INTO guests_groups (
  guests_id, guests_groups_types_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteGuestGroup :exec
DELETE FROM guests_groups
WHERE id = $1;

-- name: CreateGuestGroupType :one
INSERT INTO guests_groups_types (
  creator_id, group_name, group_color
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteGuestGroupType :exec
DELETE FROM guests_groups_types
WHERE id = $1;



