package module

import (
	"fmt"
	"github.com/gestgo/gest/package/technique/validate"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"log"
	"net/http"
	"payment/config"
	"payment/docs"
	"payment/src/echoSwagger"
	"time"
)

func EnableValidationRequest(e *echo.Echo, validator validate.IValidator) {
	e.Validator = validator

}
func EnableLogRequest(e *echo.Group) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

}

func EnableSwagger(e *echo.Group, logger *zap.SugaredLogger) {
	docs.SwaggerInfo.BasePath = config.GetConfiguration().Http.BasePath
	logger.Infof("swagger: http://0.0.0.0:%v%s/api-docs", config.GetConfiguration().Http.Port, config.GetConfiguration().Http.BasePath)
	//swaggerHandler := echoSwagger.EchoWrapHandler(echoSwagger.URL(swaggerURL))
	e.GET("/*", echoSwagger.WrapHandler)
}

func EnableErrorHandler(e *echo.Echo) {

	e.HTTPErrorHandler = customHTTPErrorHandler

}

func SetGlobalPrefix(e *echo.Echo) *echo.Group {
	return e.Group(config.GetConfiguration().Http.BasePath)
}

func EnableNotFound(e *echo.Echo, group *echo.Group) {
	e.Any("/*", customHTTP404RouterHandler)

	group.Any("/*", customHTTP404RouterHandler)

}

func customHTTPErrorHandler(err error, c echo.Context) {

	code := http.StatusInternalServerError
	//he, ok := err.(*echo.HTTPError)

	if he, ok := err.(*echo.HTTPError); ok {
		// Check if the error is an HTTPError object
		log.Print(he, ok)
		if he.Code == http.StatusBadRequest {
			// If the error code is 404, send a "path not found" message
			//message := "Path not found"
			error400 := HttpError[any]{
				StatusCode: he.Code,
				Message:    he.Message,
				Path:       c.Request().URL.Path,
				Timestamp:  time.Now().UnixMilli(),
			}
			c.JSON(http.StatusBadRequest, error400)
			return
		}
	}

	errorRes := HttpError[string]{
		StatusCode: code,
		Message:    "Internal Server Error",
		Path:       c.Request().URL.Path,
		Timestamp:  time.Now().UnixMilli(),
	}
	c.JSON(code, errorRes)
}
func customHTTP404RouterHandler(c echo.Context) error {
	code := http.StatusNotFound
	errorRes := HttpError[string]{
		StatusCode: code,
		Message:    fmt.Sprintf("Cannot %s %s", c.Request().Method, c.Request().URL.Path),
		Path:       c.Request().URL.Path,
		Timestamp:  time.Now().UnixMilli(),
	}
	return c.JSON(code, errorRes)
}

type HttpError[T any] struct {
	StatusCode int    `json:"statusCode"`
	Message    T      `json:"message"`
	Path       string `json:"path"`
	Timestamp  int64  `json:"timestamp"`
}
