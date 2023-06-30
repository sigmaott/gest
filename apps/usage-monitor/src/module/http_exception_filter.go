package module

import (
	errorGest "github.com/gestgo/gest/package/core/error"
)

type BadRequestError[T any] struct {
	errorGest.HttpError[T]
	Errors any `json:"errors,omitempty"`
}
