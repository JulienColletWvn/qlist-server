-- name: CreateEvent :one
INSERT INTO events (
  name, description, start_date, end_date, location, free_wifi, public, tickets_amount
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: UpdateEventName :one
UPDATE events
SET name = $2, updated_at = now()
WHERE id = $1
RETURNING *;

-- name: UpdateEventDescription :one
UPDATE events
SET description = $2, updated_at = now()
WHERE id = $1
RETURNING *;

-- name: UpdateEventDates :one
UPDATE events
SET start_date = $2, end_date = $3, updated_at = now()
WHERE id = $1
RETURNING *;

-- name: UpdateEventLocation :one
UPDATE events
SET location = $2, updated_at = now()
WHERE id = $1
RETURNING *;

-- name: UpdateEventWifiAvailability :one
UPDATE events
SET free_wifi = $2, updated_at = now()
WHERE id = $1
RETURNING *;

-- name: UpdateEventIsPublic :one
UPDATE events
SET public = $2, updated_at = now()
WHERE id = $1
RETURNING *;

-- name: UpdateEventTicketsAmount :one
UPDATE events
SET tickets_amount = $2, updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteEvent :exec
DELETE FROM events
WHERE id = $1;

-- name: GetUserEvents :many
SELECT * FROM events
WHERE creator_id = $1
ORDER BY start_date DESC
LIMIT $1
OFFSET $2;

-- name: GetPublicEvents :many
SELECT * FROM events
WHERE public = true
ORDER BY start_date DESC
LIMIT $1
OFFSET $2;

-- name: CreateEventAdministrator :one
INSERT INTO events_administrators (
  users_id, events_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteEventAdministrator :exec
DELETE FROM events_administrators
WHERE users_id = $1;


-- name: CreateEventPhoto :one
INSERT INTO events_photos (
  events_id, url
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteEventPhoto :exec
DELETE FROM events_photos
WHERE id = $1;

-- name: CreateEventGuest :one
INSERT INTO events_guests (
  events_id, guests_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteEventGuest :exec
DELETE FROM events_guests
WHERE id = $1;

-- name: CreateEventGuestGroup :one
INSERT INTO events_guests_groups (
  events_id, guests_groups_types
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteEventGuestGroup :exec
DELETE FROM events_guests_groups
WHERE id = $1;

