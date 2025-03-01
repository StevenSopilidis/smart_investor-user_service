package app_errors

type EmailOrPasswordAlreadyExist struct{}

func (e *EmailOrPasswordAlreadyExist) Error() string {
	return "Email or password already exist"
}
