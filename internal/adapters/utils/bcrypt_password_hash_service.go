package utils

import (
	"gitlab.com/stevensopi/smart_investor/user_service/internal/core/app_errors"
	"golang.org/x/crypto/bcrypt"
)

type BcryptPasswordHashService struct{}

func NewBcryptPasswordHashService() *BcryptPasswordHashService {
	return &BcryptPasswordHashService{}
}

func (bs *BcryptPasswordHashService) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", &app_errors.PasswordHashFailed{}
	}

	return string(hash), nil
}

func (bs *BcryptPasswordHashService) VerifyPassword(password string, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return &app_errors.InvalidPassword{}
	}

	return nil
}
