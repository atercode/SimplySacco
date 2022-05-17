package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

const minSecretKeySize = 32

// JWTMaker is a JSON Web Token Maker
type JWTMaker struct {
	secretKey string
}

// NewJWTMaker creates a new JWTMaker
func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("Invalid key size; secret key must be atleast %d characters long.", minSecretKeySize)
	}
	return &JWTMaker{secretKey: secretKey}, nil
}

//CreateToken creates a token for a specific username and duration
func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(maker.secretKey))
}

//VerifyToken checks validity of the given token
func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	//prevents encoding mismatch attacks by double checking the encoding submited in request vs encoding used on the server
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrorInvalidToken
		}
		return []byte(maker.secretKey), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrorExpiredToken) {
			return nil, ErrorExpiredToken
		}
		return nil, ErrorInvalidToken
	}
	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrorInvalidToken
	}
	return payload, nil
}
