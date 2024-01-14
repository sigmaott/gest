package interceptor

import "github.com/labstack/echo/v4"

type Interceptor func(c echo.Context, next echo.HandlerFunc) error

func UseInterceptors(interceptors ...Interceptor) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var chain func(i int, ctx echo.Context, nextChain echo.HandlerFunc) error
			chain = func(i int, ctxChain echo.Context, nextChain echo.HandlerFunc) error {
				if i < len(interceptors) {
					return interceptors[i](ctxChain, func(ctx echo.Context) error {
						return chain(i+1, ctx, nextChain)
					})
				}
				return nextChain(c)
			}

			return chain(0, c, next)
		}
	}
}
