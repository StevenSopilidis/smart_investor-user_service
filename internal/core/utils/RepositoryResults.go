package utils

const (
	EmailVerifiedSuccessfully uint8 = iota
	UserDeletedSuccessfully
	EmailOrPasswordAlreadyExists
	UserNotFound
	InternalServerError
)
