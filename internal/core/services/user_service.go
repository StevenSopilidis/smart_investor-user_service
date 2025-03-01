package services

import (
	"errors"

	"gitlab.com/stevensopi/smart_investor/user_service/internal/core/app_errors"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/core/domain"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/core/dtos"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/core/ports"
)

type UserService struct {
	repo ports.IUserRepo
}

func NewUserService(repo ports.IUserRepo) (*UserService, error) {
	return &UserService{
		repo: repo,
	}, nil
}

func (s *UserService) CreateUser(dto dtos.CreateUserDto) error {
	_, err := s.repo.FindUserByEmail(dto.Email)

	if errors.Is(err, &app_errors.UserNotFound{}) {
		return s.repo.CreateUser(domain.User{
			Email:    dto.Email,
			Password: dto.Password,
		})
	}

	if err == nil {
		return &app_errors.EmailOrPasswordAlreadyExist{}
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
