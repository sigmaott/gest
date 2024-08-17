package common

import (
	"example/src/pipe"
	"fmt"
	"reflect"

	"github.com/labstack/echo/v4"
)

type QueryTransform[T any, O any] struct {
	key string
}

func (p *QueryTransform[T, O]) Get(ctx echo.Context) O {
	value := ctx.Get(p.key).(O)
	return value
}

func (p *QueryTransform[T, O]) Transform(pipes ...pipe.TransformFunc) pipe.Pipe {
	return func(c echo.Context) error {

		// is struct
		obj := new(T)

		if reflect.TypeOf(*obj).Kind() == reflect.Struct {
			if err := (&echo.DefaultBinder{}).BindQueryParams(c, obj); err != nil {

				return NewErrorTransform(err)
			}
		} else {

			if _, err := fmt.Sscanf(c.QueryParam(p.key), "%v", obj); err != nil {
				return err
			}

		}
		var err error
		var result any = *obj
		for _, pipe := range pipes {

			if result, err = pipe(*obj, p.key); err != nil {
				return err
			} else {

			}
		}

		c.Set(p.key, result)

		return nil
	}
}

func Query[T any, O any](key string) ITransform[T, O] {
	return &QueryTransform[T, O]{
		key: key,
	}
}
