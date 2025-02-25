package services

import (
	"gitlab.com/stevensopi/smart_investor/user_service/internal/core/domain"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/core/ports"
)

type UserService struct {
	repo ports.IUserRepo
}

func (s *UserService) CreateUser(user domain.User) {

}

func (s *UserService) FindUserByEmail(user domain.User) {

}

func (s *UserService) ValidateEmail(user domain.User) {

}

func (s *UserService) DeleteUser(user domain.User) {

}
