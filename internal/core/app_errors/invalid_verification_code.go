package app_errors

type InvalidVerificationCode struct{}

func (e *InvalidVerificationCode) Error() string {
	return "Something went wrong"
}

func (e *InvalidVerificationCode) Is(target error) bool {
	_, ok := target.(*InvalidVerificationCode)
	return ok
}
