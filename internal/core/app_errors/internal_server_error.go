package app_errors

type InternalServerError struct{}

func (e *InternalServerError) Error() string {
	return "Something went wrong"
}
