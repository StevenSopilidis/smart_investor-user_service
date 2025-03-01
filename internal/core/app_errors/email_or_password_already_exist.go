package app_errors

type EmailOrPasswordAlreadyExist struct{}

func (e *EmailOrPasswordAlreadyExist) Error() string {
	return "Email or password already exist"
}

func (e *EmailOrPasswordAlreadyExist) Is(target error) bool {
	_, ok := target.(*EmailOrPasswordAlreadyExist)
	return ok
}
