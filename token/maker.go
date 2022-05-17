package token

import "time"

type Maker interface {
	//CreateToken creates a token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)
	//VerifyToken checks validity of the given token
	VerifyToken(token string) (*Payload, error)
}
