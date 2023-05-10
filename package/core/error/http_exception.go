package error

type HTTPException[T any] struct {
	StatusCode   int    `json:"statusCode"`
	Message      T      `json:"message"`
	ErrorMessage string `json:"error"`
}

func (H *HTTPException[T]) Error() string {
	return H.ErrorMessage
}

func NewHTTPException[T any](statusCode int, message T, error string) error {

	return &HTTPException[T]{
		StatusCode:   statusCode,
		Message:      message,
		ErrorMessage: error,
	}

}
