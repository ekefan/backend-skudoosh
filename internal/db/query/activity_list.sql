-- name: CreateActivity :one
INSERT INTO activity_lists(
    trip_owner,
    activity,
    date_time
) VALUES (
    $1, $2, $3
) RETURNING *;


-- name: GetActivityUpdate :one
SELECT FROM activity_lists
WHERE trip_owner = $1
LIMIT 1
FOR NO KEY UPDATE;

-- name: ListActivities :many
SELECT FROM activity_lists
WHERE trip_owner = $1
LIMIT $2
OFFSET $3;

-- name: DeleteActivity :exec
DELETE FROM activity_lists WHERE trip_owner = &1;