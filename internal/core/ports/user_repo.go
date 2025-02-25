package ports

import "gitlab.com/stevensopi/smart_investor/user_service/internal/core/domain"

type IUserRepo interface {
	CreateUser(user domain.User) uint8
	FindUserByEmail(email string) uint8
	ValidateEmail(user domain.User) uint8
	DeleteUser(user domain.User) uint8
}
