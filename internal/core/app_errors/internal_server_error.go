package app_errors

type InternalServerError struct{}

func (e *InternalServerError) Error() string {
	return "Something went wrong"
}

func (e *InternalServerError) Is(target error) bool {
	_, ok := target.(*InternalServerError)
	return ok
}
