package postgres

import (
	"gitlab.com/stevensopi/smart_investor/user_service/internal/adapters/config"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/core/app_errors"
	"gitlab.com/stevensopi/smart_investor/user_service/internal/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresRepo struct {
	db *gorm.DB
}

func NewPostgresRepo(config config.Config) (*PostgresRepo, error) {
	db, err := gorm.Open(postgres.Open(config.DBConnection), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		return nil, err
	}

	return &PostgresRepo{
		db: db,
	}, nil
}

func (r *PostgresRepo) CreateUser(user domain.User) error {
	result := r.db.Create(user)
	if result.Error != nil {
		return &app_errors.InternalServerError{}
	}
	return nil
}

func (r *PostgresRepo) FindUserByEmail(email string) (domain.User, error) {
	var user domain.User
	result := r.db.Where("email=?", email).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return domain.User{}, &app_errors.UserNotFound{}
		}
		return domain.User{}, &app_errors.InternalServerError{}
	}
	return user, nil
}

func (r *PostgresRepo) ValidateEmail(user domain.User) error {
	result := r.db.Model(&user).Update("email_verified", true)
	if result.Error != nil {
		return &app_errors.InternalServerError{}
	}
	return nil
}

func (r *PostgresRepo) DeleteUser(user domain.User) error {
	result := r.db.Delete(user)
	if result.Error != nil {
		return &app_errors.InternalServerError{}
	}
	return nil
}
