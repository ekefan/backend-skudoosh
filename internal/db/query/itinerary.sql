-- name: CreateTrip :one
INSERT INTO itineraries(
    owner,
    take_off_date,
    return_date,
    destination
) VALUES (
    $1, $2, $3, $4
) RETURNING *;


-- name: GetTripUpdate :one
SELECT FROM itineraries
WHERE owner = $1
LIMIT 1
FOR NO KEY UPDATE;

-- name: ListTrips :many
SELECT FROM itineraries
WHERE owner = $1
LIMIT $2
OFFSET $3;

-- name: DeleteItinerary :exec
DELETE FROM itineraries WHERE owner = $1;