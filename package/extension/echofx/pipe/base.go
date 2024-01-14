package pipe

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type Pipe func(c echo.Context) error
type ITransformPipe interface {
	// Get(ctx echo.Context) T
	Transform(pipes ...TransformFunc) Pipe
}

// type intOrString interface {
// 	Pipe | ITransformPipe
// }

type TransformFunc func(value any, key string) (r any, e error)

func UsePipes(pipes ...any) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			for _, pipe := range pipes {
				switch p := pipe.(type) {
				case Pipe:
					if err := p(c); err != nil {
						return err
					}
					continue
				case ITransformPipe:
					if err := p.Transform()(c); err != nil {
						return err
					}
					continue
				default:
					fmt.Println("Unknown instance type")
				}
			}
			return next(c)

		}
	}
}
