package gapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/adapters/handlers/grpc/generated"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestVerifyValidEmailReturnsOK(t *testing.T) {
	res, _ := createValidUser(t)

	req := &generated.VerifyEmailRequest{
		Email:                  res.Email,
		EmailVerificiationCode: res.EmailVerificationCode,
	}

	_, err := server.VerifyEmail(context.Background(), req)
	require.NoError(t, err)
	status, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.OK, status.Code())

	user, err := server.user_service.FindUserByEmail(res.Email)
	require.NoError(t, err)
	require.True(t, user.EmailVerified)
}

func TestVerifyInvalidEmailThatDoesNotExistReturnsNotFound(t *testing.T) {
	req := &generated.VerifyEmailRequest{
		Email:                  "invalid@test.com",
		EmailVerificiationCode: "code",
	}

	_, err := server.VerifyEmail(context.Background(), req)
	require.Error(t, err)
	status, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.NotFound, status.Code())
}

func TestVerifyEmailWithInvalidCodeReturnsInvalidArgument(t *testing.T) {
	res, _ := createValidUser(t)

	req := &generated.VerifyEmailRequest{
		Email:                  res.Email,
		EmailVerificiationCode: "invalid_code",
	}

	_, err := server.VerifyEmail(context.Background(), req)
	require.Error(t, err)
	status, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.InvalidArgument, status.Code())
}
