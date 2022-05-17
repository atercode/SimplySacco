package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

//Different types of error returned by JWT Valid funtion
var (
	ErrorInvalidToken = errors.New("token is invalid")
	ErrorExpiredToken = errors.New("Token has expired")
	ErrorTimeMismatch = errors.New("Token issued at a future date")
)

//Payload contains the payload data of the token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

//NewPayload creates a new token payload with a specific username and duration
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

// Valid is an interfaced enfored class of the JWT go package that servers to check if a token payload is valid or not
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrorExpiredToken
	}
	if time.Now().Before(payload.IssuedAt) {
		return ErrorTimeMismatch
	}
	return nil
}
