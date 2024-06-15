package token

import "time"

// Maker is an interface for managing tokens
type Maker interface {
	// CreateToken creates a new token for a specific username and duration
	CreateToken(user_id int64, duration time.Duration) (string, error)

	// VerifyToken checks if the token is valid or not,
	// Returns the Payload data on true and error on false
	VerifyToken(token string) (*Payload, error)
}
