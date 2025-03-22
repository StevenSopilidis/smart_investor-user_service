package app_errors

type PasswordHashFailed struct{}

func (e *PasswordHashFailed) Error() string {
	return "Could not hash password"
}

func (e *PasswordHashFailed) Is(target error) bool {
	_, ok := target.(*PasswordHashFailed)
	return ok
}
