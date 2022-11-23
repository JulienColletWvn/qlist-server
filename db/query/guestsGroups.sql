-- name: CreateGuestsGroupType :one
INSERT INTO guests_groups_types (
        creator_id,
        events_id,
        group_name,
        group_color
    )
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: DeleteGuestsGroupType :exec
DELETE FROM guests_groups_types
WHERE guests_groups_types.id = $1
    AND guests_groups_types.events_id IN (
        SELECT events_administrators.events.id
        FROM events_administrators
        WHERE events_administrators.users_id = $2
    );
-- name: GetGuestsGroupType :one
SELECT *
FROM guests_groups_types
WHERE guests_groups_types.id = $1
    AND guests_groups_types.events_id IN (
        SELECT events_administrators.events.id
        FROM events_administrators
        WHERE events_administrators.events_id = $2
            AND events_administrators.users_id = $3
    );
-- name: GetGuestsGroupTypes :one
SELECT *
FROM guests_groups_types
WHERE guests_groups_types.events_id IN (
        SELECT events_administrators.events.id
        FROM events_administrators
        WHERE events_administrators.events_id = $1
            AND events_administrators.users_id = $2
    );
-- name: CreateGuestsGroup :one
INSERT INTO guests_groups (guests_id, guests_groups_types_id)
VALUES ($1, $2)
RETURNING *;
-- name: DeleteGuestsGroup :exec
DELETE FROM guests_groups g USING guests_groups_types t
WHERE g.id = $1
    AND g.guests_groups_types_id = t.id
    AND t.events_id IN (
        SELECT events_administrators.events.id
        FROM events_administrators
        WHERE events_administrators.users_id = $2
    );
-- name: GetGuestsGroup :one
SELECT *
FROM guests_groups g
    LEFT JOIN guests_groups_types t ON t.id = g.guests_groups_types_id
WHERE g.id = $1
    AND t.events_id IN (
        SELECT events_administrators.events.id
        FROM events_administrators
        WHERE events_administrators.events_id = $2
            AND events_administrators.users_id = $3
    );
-- name: GetGuestsGroups :one
SELECT *
FROM guests_groups g
    LEFT JOIN guests_groups_types t ON t.id = g.guests_groups_types_id
WHERE t.events_id IN (
        SELECT events_administrators.events.id
        FROM events_administrators
        WHERE events_administrators.events_id = $1
            AND events_administrators.users_id = $2
    );