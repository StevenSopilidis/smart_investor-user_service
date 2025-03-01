package services

import (
	"errors"

	"gitlab.com/stevensopi/smart_investor/user_service/internal/core/app_errors"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/core/domain"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/core/ports"
)

type UserService struct {
	repo ports.IUserRepo
}

func (s *UserService) CreateUser(user domain.User) error {
	user, err := s.repo.FindUserByEmail(user.Email)

	if errors.Is(err, &app_errors.UserNotFound{}) {
		return s.repo.CreateUser(user)
	}

	return err
}

func (s *UserService) FindUserByEmail(email string) (domain.User, error) {
	return s.repo.FindUserByEmail(email)
}

func (s *UserService) ValidateEmail(user domain.User) error {
	return s.repo.ValidateEmail(user)
}

func (s *UserService) DeleteUser(user domain.User) error {
	return s.repo.DeleteUser(user)
}
