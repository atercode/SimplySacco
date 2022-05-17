package utils

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestPassword(t *testing.T) {
	password := gofakeit.Password(true, true, true, true, false, 8)

	hashedPassword, err := HashPassword(password)

	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = CheckPassword(password, hashedPassword)
	require.NoError(t, err)

	wrongPassword := gofakeit.Password(true, true, true, true, false, 6)
	err = CheckPassword(wrongPassword, hashedPassword)
	require.Error(t, err)
}
