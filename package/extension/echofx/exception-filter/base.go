package exception_filter

import "github.com/labstack/echo/v4"

type ExceptionFilter func(err error, c echo.Context) error

func UseFilters(filters ...ExceptionFilter) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err != nil {
				for _, filter := range filters {
					if err = filter(err, c); err == nil {
						return err
					}
				}
			}

			return err
		}
	}
}
