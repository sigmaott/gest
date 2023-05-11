package parser

import (
	"github.com/gestgo/gest/package/core/pipe"
	"github.com/labstack/echo/v4"
)

type BindBody interface {
	BindBody(c echo.Context, i interface{}) error
}
type BodyParser[T any] struct {
	name     string
	validate bool
	binder   BindBody
	pipes    []pipe.IPipe
}

func (b *BodyParser[T]) Parser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		data := new(T)
		err := b.binder.BindBody(c, data)
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

func NewBodyParser[T any](name string, validate bool, p ...pipe.IPipe) IParser {
	return &BodyParser[T]{
		name:     name,
		binder:   &echo.DefaultBinder{},
		validate: validate,
		pipes:    p,
	}
}
