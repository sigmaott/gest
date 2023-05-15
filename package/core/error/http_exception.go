package error

type HttpError[T any] struct {
	StatusCode int    `json:"statusCode"`
	Message    T      `json:"message"`
	Path       string `json:"path"`
	Timestamp  int64  `json:"timestamp"`
}
