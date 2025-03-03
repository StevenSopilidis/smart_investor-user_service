package ports

import "gitlab.com/stevensopi/smart_investor/user_service/internal/core/domain"

type IUserRepo interface {
	CreateUser(user domain.User) error
	FindUserByEmail(email string) (domain.User, error)
	ValidateEmail(user domain.User) error
	DeleteUser(user domain.User) error
}
