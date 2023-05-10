package exceptions

import (
	"github.com/labstack/echo/v4"
)

type IEchoCustomException interface {
	ErrorHandler(err error, c echo.Context)
}
