package token

import (
	"time"

	"github.com/vk-rv/pvx"
)

type PasetoMaker struct {
	paseto       *pvx.ProtoV4Local
	symmetricKey *pvx.SymKey
}

// NewPasetoMaker creates a new Paseto Maker instance
func NewPasetoMaker(keyMaterial string) (Maker, error) {

	symmetricKey := pvx.NewSymmetricKey([]byte(keyMaterial), pvx.Version4)
	paseto := pvx.NewPV4Local()
	maker := &PasetoMaker{
		paseto:       paseto,
		symmetricKey: symmetricKey,
	}
	return maker, nil
}

//CreateToken creates a token for a specific username and duration
func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", nil
	}
	return maker.paseto.Encrypt(maker.symmetricKey, payload, pvx.WithAssert([]byte("test")))
}

//VerifyToken checks validity of the given token
func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, pvx.WithAssert([]byte("test"))).ScanClaims(payload)

	if err != nil {
		return nil, err
	}
	err = payload.Valid()
	if err != nil {
		return nil, err
	}
	return payload, nil

}
