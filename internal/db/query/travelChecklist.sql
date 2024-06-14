-- name: CreateTravelChecklist :one
INSERT INTO travel_checklists(
    trip_owner,
    item_task
) VALUES (
    $1, $2
) RETURNING *;


-- name: GetTravelChecklist :one
SELECT FROM travel_checklists
WHERE trip_owner = $1
LIMIT 1
FOR NO KEY UPDATE;

-- name: ListTravelChecklist :many
SELECT FROM travel_checklists
WHERE trip_owner = $1
LIMIT $2
OFFSET $3;

-- name: DeleteTravelChecklist :exec
DELETE FROM travel_checklists WHERE trip_owner = &1;