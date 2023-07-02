package filter

import "github.com/labstack/echo/v4"

type EchoErrorFilter struct {
	filters []func(err error, c echo.Context) (code int, res any)
}

func (e *EchoErrorFilter) CatchError(err error, c echo.Context) {
	for _, filter := range e.filters {
		code, res := filter(err, c)
		if code != 0 {
			c.JSON(code, res)
			return
		}

	}

}
func NewEchoExceptionFilter(filters ...func(err error, c echo.Context) (code int, res any)) *EchoErrorFilter {
	return &EchoErrorFilter{
		filters: filters,
	}

}
