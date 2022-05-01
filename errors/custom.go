package errors

type CustomError struct {
	msg string
}

func (e *CustomError) Error() string {
	return e.msg
}

func NewCustomError(msg string) error {
	return &CustomError{msg}
}

func (e *CustomError) Is(target error) bool {
	t, ok := target.(*CustomError)

	if !ok {
		return false
	}

	return e.Error() == t.Error()
}
