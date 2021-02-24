package token

import "time"

// Maker in an interface for managing tokens
type Maker interface {
	// CreateToken creates a new token for a specific username as duration
	CreateToken(username string, duration time.Duration) (string, error)

	//  VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
