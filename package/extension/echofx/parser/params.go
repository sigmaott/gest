package parser

import (
	"github.com/gestgo/gest/package/core/pipe"
	"github.com/labstack/echo/v4"
)

type BindPathParams interface {
	BindPathParams(c echo.Context, i interface{}) error
}
type PathParams[T any] struct {
	name     string
	validate bool
	binder   BindPathParams
	pipes    []pipe.IPipe
}

func (b *PathParams[T]) Parser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		data := new(T)
		err := b.binder.BindPathParams(c, data)

		if err != nil {
			return c.JSON(400, err)
		}
		if b.validate {
			if err = c.Validate(data); err != nil {
				return err
			}

		}
		for _, pipe := range b.pipes {
			if err = pipe.Bind(data); err != nil {
				return err
			}

		}
		c.Set(b.name, data)

		err = next(c)

		return err
	}
}

func NewParamsParser[T any](name string, validate bool, p ...pipe.IPipe) IParser {
	return &PathParams[T]{
		name:     name,
		binder:   &echo.DefaultBinder{},
		validate: validate,
		pipes:    p,
	}
}
