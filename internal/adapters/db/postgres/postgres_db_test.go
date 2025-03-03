package postgres

import (
	"crypto/rand"
	"log"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/adapters/config"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/core/app_errors"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/core/domain"
)

func generateRandomString(n int) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var sb strings.Builder
	for i := 0; i < 10; i++ {
		randNum, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		sb.WriteByte(letters[randNum.Int64()])
	}

	return sb.String()
}

func generateRandomEmail() string {
	return generateRandomString(10) + "@test.com"
}

func generateRandomPassword() string {
	return generateRandomString(15)
}

var testRepo *PostgresRepo

func TestMain(m *testing.M) {
	config := config.Config{
		DBConnection: "postgresql://root:secret@localhost:5432/smart_investor_user_test_db?sslmode=disable",
	}

	repo, err := NewPostgresRepo(config)
	if err != nil {
		log.Fatal("Could not create test repo: ", err)
	}

	testRepo = repo

	code := m.Run()
	os.Exit(code)
}

func createUserAndTest(t *testing.T) domain.User {
	user := domain.User{
		Id:                    uuid.New(),
		Email:                 generateRandomEmail(),
		Password:              generateRandomPassword(),
		EmailVerified:         false,
		EmailVerificationCode: generateRandomString(20),
		CreatedAt:             time.Now(),
	}

	err := testRepo.CreateUser(user)
	require.NoError(t, err)

	result, err := testRepo.FindUserByEmail(user.Email)
	require.NoError(t, err)
	require.Equal(t, user.Id, result.Id)
	require.Equal(t, user.Email, result.Email)
	require.Equal(t, user.Password, result.Password)
	require.Equal(t, user.EmailVerificationCode, result.EmailVerificationCode)
	require.WithinDuration(t, user.CreatedAt, result.CreatedAt, time.Second)
	require.Equal(t, user.EmailVerified, result.EmailVerified)

	return user
}

func TestCreateUserPostgresRepo(t *testing.T) {
	createUserAndTest(t)
}

func TestValidateEmailPostgresRepo(t *testing.T) {
	user := createUserAndTest(t)

	err := testRepo.ValidateEmail(user)
	require.NoError(t, err)

	updatedUser, err := testRepo.FindUserByEmail(user.Email)
	require.NoError(t, err)
	require.True(t, updatedUser.EmailVerified)
}

func TestDeleteUserPostgreRepo(t *testing.T) {
	user := createUserAndTest(t)

	err := testRepo.DeleteUser(user)
	require.NoError(t, err)

	user, err = testRepo.FindUserByEmail(user.Email)
	require.ErrorIs(t, err, &app_errors.UserNotFound{})
}
