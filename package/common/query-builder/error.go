package query_builder

type QueryParserError struct {
	Message string
}

func (q QueryParserError) Error() string {
	return q.Message
}

func NewQueryParserError(message string) error {
	return &QueryParserError{
		Message: message,
	}
}

type QueryParserErrors []QueryParserError

func (q QueryParserErrors) Error() string {
	return "have error"
}

func NewQueryParsersError(message string) error {
	return &QueryParserErrors{}
}

type ValidateError struct {
	Message string
}

func (v ValidateError) Error() string {
	return v.Message
}

func NewValidateError(message string) error {
	return &ValidateError{
		Message: message,
	}
}
