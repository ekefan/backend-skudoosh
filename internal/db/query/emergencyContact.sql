-- name: CreateEmergencyContact :one
INSERT INTO emergency_contacts(
    owner,
    email,
    phone_number
) VALUES (
    $1, $2, $3
) RETURNING *;


-- name: GetEmergencyContactUpdate :one
SELECT * FROM emergency_contacts
WHERE owner = $1
LIMIT 1
FOR NO KEY UPDATE;

-- name: ListsAccounts :many
SELECT * FROM emergency_contacts
WHERE owner = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- DeleteEmergencyContact :exec
DELETE FROM emergency_contacts WHERE owner = $1;

