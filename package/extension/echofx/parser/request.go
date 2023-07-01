package parser

import (
	"github.com/gestgo/gest/package/core/pipe"
	"github.com/labstack/echo/v4"
)

type BindRequest interface {
	Bind(i interface{}, c echo.Context) error
}
type RequestParams[T any] struct {
	name     string
	validate bool
	binder   BindRequest
	pipes    []pipe.IPipe
}

func (b *RequestParams[T]) Parser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		data := new(T)
		err := b.binder.Bind(data, c)
		if err != nil {
			return err
		}
		if b.validate {
			if err = c.Validate(data); err != nil {
				return err
			}

		}
		c.Set(b.name, data)

		err = next(c)

		return err
	}
}

func NewRequestParser[T any](name string, validate bool, p ...pipe.IPipe) IParser {
	return &RequestParams[T]{
		name:     name,
		binder:   &echo.DefaultBinder{},
		validate: validate,
		pipes:    p,
	}
}
