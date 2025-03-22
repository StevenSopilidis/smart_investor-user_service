package gapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/adapters/handlers/grpc/generated"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestFindValidUserReturnsOk(t *testing.T) {
	res, _ := createValidUser(t)

	user, err := server.FindUserByEmail(context.Background(), &generated.FindUserByEmailRequest{
		Email: res.Email,
	})
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, res.Email, user.Email)
	require.Equal(t, res.EmailVerified, user.EmailVerified)
	require.Equal(t, res.Id, user.Id)
}

func TestFindInvalidUserReturnsNotFound(t *testing.T) {
	email := generateRandomEmail()

	_, err := server.FindUserByEmail(context.Background(), &generated.FindUserByEmailRequest{
		Email: email,
	})
	require.Error(t, err)
	status, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.NotFound, status)
}
