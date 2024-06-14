-- name: CreateTripLogs :one
INSERT INTO trip_logs(
    trip_owner,
    logs
) VALUES (
    $1, $2
) RETURNING *;


-- name: GetATripLogUpdate :one
SELECT FROM trip_logs
WHERE trip_owner = $1
LIMIT 1
FOR NO KEY UPDATE;

-- name: ListTripLogs :many
SELECT FROM trip_logs
WHERE trip_owner = $1
LIMIT $2
OFFSET $3;

-- name: DeleteTripLogs :exec
DELETE FROM trip_logs WHERE trip_owner = &1;