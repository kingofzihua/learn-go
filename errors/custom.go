package errors

type CustomError struct {
	msg string
}

func (e CustomError) Error() string {
	return e.msg
}

func NewCustomError(msg string) error {
	return CustomError{msg}
}
