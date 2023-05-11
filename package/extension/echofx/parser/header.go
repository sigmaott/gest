package parser

import (
	"github.com/gestgo/gest/package/core/pipe"
	"github.com/labstack/echo/v4"
)

type BindHeaders interface {
	BindHeaders(c echo.Context, i interface{}) error
}
type HeaderParser[T any] struct {
	name     string
	binder   BindHeaders
	validate bool
	pipes    []pipe.IPipe
}

func (b *HeaderParser[T]) Parser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		data := new(T)
		err := b.binder.BindHeaders(c, data)
		if err != nil {
			return err
		}

		for _, pipe := range b.pipes {
			if err = pipe.Bind(data); err != nil {
				return err
			}

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

func NewDefaultHeaderParser[T any](name string, validate bool, p ...pipe.IPipe) IParser {
	return &HeaderParser[T]{
		name:     name,
		binder:   &echo.DefaultBinder{},
		validate: validate,
		pipes:    p,
	}
}
