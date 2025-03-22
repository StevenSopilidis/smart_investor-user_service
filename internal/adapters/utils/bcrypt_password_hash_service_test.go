package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/core/app_errors"
)

func TestVerifyValidPAsswordReturnsNoError(t *testing.T) {
	password := "password123"

	service := NewBcryptPasswordHashService()
	hash, err := service.HashPassword(password)
	require.NoError(t, err)
	require.NoError(t, service.VerifyPassword(password, hash))
}

func TestVerifyInvalidPasswordReturnsInvalidPasswordError(t *testing.T) {
	password := "password123"

	service := NewBcryptPasswordHashService()
	hash, err := service.HashPassword(password)
	require.NoError(t, err)
	err = service.VerifyPassword("invalid", hash)
	require.ErrorIs(t, err, &app_errors.InvalidPassword{})
}
