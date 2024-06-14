-- name: CreateTripBooking :one
INSERT INTO trip_bookings(
    trip_owner,
    booking_type,
    booking_details
) VALUES (
    $1, $2, $3
) RETURNING *;


-- name: GetTripBookingUpdate :one
SELECT * FROM trip_bookings
WHERE trip_owner = $1
LIMIT 1
FOR NO KEY UPDATE;

-- name: ListTripBooking :many
SELECT FROM trip_bookings
WHERE trip_owner = $1
LIMIT $2
OFFSET $3;

-- name: DeleteTripBooking :exec
DELETE FROM trip_bookings WHERE trip_owner = &1;