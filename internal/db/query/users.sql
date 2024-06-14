-- name: CreateUser :one
INSERT INTO users (
    fullname,
    username, 
    email, 
    hashed_password,
    phone_number
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users 
WHERE username = $1
LIMIT 1;

-- name: GetUserForUpdate :one
SELECT * FROM users
WHERE username = $1
LIMIT 1
FOR NO KEY UPDATE;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = &1;