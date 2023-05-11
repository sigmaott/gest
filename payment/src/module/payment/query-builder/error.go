package query_builder

type ValidateError struct {
	error error
}

func (i *ValidateError) Error() string {
	return i.error.Error()
}

func NewValidateError(err error) error {
	return &ValidateError{
		error: err,
	}
}
