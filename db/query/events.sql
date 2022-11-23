-- name: CreateEvent :one
INSERT INTO events (
    start_date,
    end_date,
    location,
    free_wifi,
    public,
    tickets_amount,
    creator_id
  )
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;
-- name: DeleteEvent :exec
DELETE FROM events
WHERE events.id = $1
  AND events.creator_id IN (
    SELECT events_administrators.users_id
    FROM events_administrators
    WHERE events_administrators.events_id = $1
  );
-- name: GetAdministratorEvent :one
SELECT *
FROM events
WHERE $1::int IN (
    SELECT events_administrators.users_id
    FROM events_administrators
    WHERE events_administrators.events_id = $2
  );
-- name: GetAdministratorEvents :many
SELECT *
FROM events
WHERE $1::int IN (
    SELECT users_id
    FROM events_administrators
  );
-- name: GetEventContents :many
SELECT *
FROM events_contents
WHERE events_id = $1;
-- name: GetPublicEvent :one
SELECT *
FROM events
WHERE id = $1
  AND events.public = true;
-- name: GetPublicEvents :many
SELECT *
FROM events
WHERE events.public = true;
-- name: CreateEventAdministrator :one
INSERT INTO events_administrators (users_id, events_id)
VALUES ($1, $2)
RETURNING *;
-- name: CreateEventContent :one
INSERT INTO events_contents (type, content, lang, events_id)
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: CreateEventPhoto :one
INSERT INTO events_photos (url, events_id)
VALUES ($1, $2)
RETURNING *;