package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/core/app_errors"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/core/domain"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/core/dtos"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/core/ports"
)

type UserService struct {
	repo                        ports.IUserRepo
	stringGenerator             ports.IRandomStringGenerator
	emailVerificationCodeLength uint8
}

func NewUserService(
	repo ports.IUserRepo,
	stringGenerator ports.IRandomStringGenerator,
	emailVerificationCodeLength uint8,
) (*UserService, error) {

	return &UserService{
		repo:                        repo,
		stringGenerator:             stringGenerator,
		emailVerificationCodeLength: emailVerificationCodeLength,
	}, nil
}

func (s *UserService) CreateUser(dto dtos.CreateUserDto) error {
	_, err := s.repo.FindUserByEmail(dto.Email)

	if errors.Is(err, &app_errors.UserNotFound{}) {
		return s.repo.CreateUser(domain.User{
			Id:                    uuid.New(),
			Email:                 dto.Email,
			Password:              dto.Password,
			CreatedAt:             time.Now(),
			EmailVerified:         false,
			EmailVerificationCode: s.stringGenerator.Generate(int(s.emailVerificationCodeLength)),
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

func (s *UserService) ValidateEmail(email string, verificationCode string) error {
	user, err := s.repo.FindUserByEmail(email)

	if err != nil {
		return err
	}

	return s.repo.ValidateEmail(user, verificationCode)
}

func (s *UserService) DeleteUser(email string) error {
	user, err := s.repo.FindUserByEmail(email)

	if err != nil {
		return err
	}

	return s.repo.DeleteUser(user)
}
