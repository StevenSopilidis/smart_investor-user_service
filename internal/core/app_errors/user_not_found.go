package app_errors

type UserNotFound struct{}

func (e *UserNotFound) Error() string {
	return "User was not found"
}
