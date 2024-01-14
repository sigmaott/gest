package common

import (
	"fmt"
	"reflect"

	"github.com/sigmaott/gest/package/extension/echofx/pipe"

	"github.com/labstack/echo/v4"
)

type ParamTransform[T any, O any] struct {
	key string
}

func (p *ParamTransform[T, O]) Get(ctx echo.Context) O {
	value := ctx.Get(p.key).(O)
	return value
}

func (p *ParamTransform[T, O]) Transform(pipes ...pipe.TransformFunc) pipe.Pipe {
	return func(c echo.Context) error {
		// is struct
		obj := new(T)
		if reflect.TypeOf(*obj).Kind() == reflect.Struct {
			if err := (&echo.DefaultBinder{}).BindPathParams(c, obj); err != nil {

				return NewErrorTransform(err)
			}
		} else {

			if _, err := fmt.Sscanf(c.Param(p.key), "%v", obj); err != nil {

				return NewErrorTransform(err)
			}

		}
		var err error
		var result any = *obj
		for _, pipe := range pipes {

			if result, err = pipe(*obj, p.key); err != nil {
				return err
			}
		}

		c.Set(p.key, result)

		return nil
	}
}

func Param[T any, O any](key string) ITransform[T, O] {
	return &ParamTransform[T, O]{
		key: key,
	}
}
