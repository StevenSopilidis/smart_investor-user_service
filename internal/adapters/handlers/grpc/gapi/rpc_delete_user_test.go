package gapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/adapters/handlers/grpc/generated"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestValidUserReturnsOK(t *testing.T) {
	res, _ := createValidUser(t)

	_, err := server.DeleteUser(context.Background(), &generated.DeleteUserRequest{
		Email: res.Email,
	})

	require.NoError(t, err)
	status, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.OK, status.Code())
}

func TestInvalidUserReturnsNotFound(t *testing.T) {
	_, err := server.DeleteUser(context.Background(), &generated.DeleteUserRequest{
		Email: "invalid@test.com",
	})

	require.Error(t, err)
	status, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.NotFound, status.Code())
}
