package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id                    uuid.UUID `gorm:primaryKey`
	Email                 string
	EmailVerified         bool
	EmailVerificationCode string
	Password              string
	CreatedAt             time.Time
}
