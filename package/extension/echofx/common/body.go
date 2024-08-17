package common

import (
	"example/src/pipe"
	"fmt"
	"reflect"

	"github.com/labstack/echo/v4"
)

type BodyTransform[T any, O any] struct {
	key string
}

func (p *BodyTransform[T, O]) Get(ctx echo.Context) O {
	value := ctx.Get(p.key).(O)
	return value
}

func (p *BodyTransform[T, O]) Transform(pipes ...pipe.TransformFunc) pipe.Pipe {
	return func(c echo.Context) error {
		obj := new(T)

		if reflect.TypeOf(*obj).Kind() == reflect.Struct {
			if err := (&echo.DefaultBinder{}).BindBody(c, obj); err != nil {
				return NewErrorTransform(err)
			}
		} else {

			if _, err := fmt.Sscanf(c.FormValue(p.key), "%v", obj); err != nil {
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

func Body[T any, O any](key string) ITransform[T, O] {
	return &BodyTransform[T, O]{
		key: key,
	}
}
