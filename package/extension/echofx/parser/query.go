package parser

import (
	"github.com/gestgo/gest/package/core/pipe"
	"github.com/labstack/echo/v4"
)

type BindQueryParams interface {
	BindQueryParams(c echo.Context, i interface{}) error
}
type QueryParams[T any] struct {
	name     string
	validate bool
	binder   BindQueryParams
	pipes    []pipe.IPipe
}

func (b *QueryParams[T]) Parser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		data := new(T)
		err := b.binder.BindQueryParams(c, data)
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

func NewQueryParser[T any](name string, validate bool, p ...pipe.IPipe) IParser {
	return &QueryParams[T]{
		name:     name,
		binder:   &echo.DefaultBinder{},
		validate: validate,
		pipes:    p,
	}
}
