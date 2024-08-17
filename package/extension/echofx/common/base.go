package common

import (
	"github.com/sigmaott/gest/package/extension/echofx/pipe"

	"github.com/labstack/echo/v4"
)

type ITransform[T any, O any] interface {
	Get(ctx echo.Context) O
	Transform(pipes ...pipe.TransformFunc) pipe.Pipe
}

// type TransformFunc func(value any, key string) (r any, e error)

func ConvertAnyNumberToFive() pipe.TransformFunc {

	return func(value any, key string) (r any, e error) {

		return 5, nil
	}

}

type ErrorTransform struct {
	error
}

func NewErrorTransform(e error) error {

	return &ErrorTransform{
		e,
	}

}
