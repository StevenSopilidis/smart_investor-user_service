package gapi

import (
	"context"
	"crypto/rand"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/adapters/handlers/grpc/generated"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func generateRandomString(n int) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var sb strings.Builder
	for i := 0; i < n; i++ {
		randNum, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		sb.WriteByte(letters[randNum.Int64()])
	}

	return sb.String()
}

func generateRandomEmail() string {
	return generateRandomString(10) + "@test.com"
}

func generateRandomPassword(n int) string {
	return generateRandomString(n)
}

func TestCreateValidUserReturnsStatusOK(t *testing.T) {
	createValidUser(t)
}

func TestCreateUserWithShortPasswordReturnsStatusInvalidArguments(t *testing.T) {
	email := generateRandomEmail()
	password := generateRandomPassword(3)

	req := &generated.CreateUserRequest{
		Email:    email,
		Password: password,
	}

	_, err := server.CreateUser(context.Background(), req)
	require.Error(t, err)
	status, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.InvalidArgument, status.Code())
}

func TestCreateUserWithInvalidEmailReturnsStatusInvalidArguments(t *testing.T) {
	email := "invalid"
	password := generateRandomPassword(15)

	req := &generated.CreateUserRequest{
		Email:    email,
		Password: password,
	}

	_, err := server.CreateUser(context.Background(), req)
	require.Error(t, err)
	status, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.InvalidArgument, status.Code())
}

func TestCreateUserThatAlreadyExistsReturnsStatusAlreadyExists(t *testing.T) {
	res, password := createValidUser(t)

	req := &generated.CreateUserRequest{
		Email:    res.Email,
		Password: password,
	}

	_, err := server.CreateUser(context.Background(), req)
	require.Error(t, err)
	status, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.AlreadyExists, status.Code())
}

func createValidUser(t *testing.T) (*generated.CreateUserResponse, string) {
	email := generateRandomEmail()
	password := generateRandomPassword(12)

	req := &generated.CreateUserRequest{
		Email:    email,
		Password: password,
	}

	res, err := server.CreateUser(context.Background(), req)
	require.NoError(t, err)
	status, ok := status.FromError(err)
	require.True(t, ok)
	require.Equal(t, codes.OK, status.Code())

	require.Equal(t, email, res.Email)
	require.False(t, res.EmailVerified)
	require.NotEmpty(t, res.EmailVerificationCode)
	require.WithinDuration(t, time.Now(), res.CreatedAt.AsTime(), time.Second)

	return res, password
}
