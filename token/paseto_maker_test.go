package token

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(gofakeit.DigitN(32))
	require.NoError(t, err)
	require.NotEmpty(t, maker)

	username := gofakeit.Username()
	duration := time.Minute
	// issuedAt := time.Now()
	// expiredAt := time.Now().Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.Equal(t, username, payload.Username)
}
func TestExpiredPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(gofakeit.DigitN(32))
	require.NoError(t, err)

	username := gofakeit.Username()
	duration := -time.Minute

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrorExpiredToken.Error())
	require.Nil(t, payload)
}
