package ports

type IPasswordHashService interface {
	HashPassword(password string) (string, error)
	VerifyPassword(password string, hash string) error
}
