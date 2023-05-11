package module

import (
	"fmt"
	"github.com/gestgo/gest/package/technique/validate"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"net/http"
	"payment/config"
	"payment/docs"
	"payment/src/echoSwagger"
	"time"
)

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
	echoExceptionFilter := NewEchoExceptionFilter(BadRequestErrorFilter, ValidateErrorFilter, InternalServerErrorFilter)
	e.HTTPErrorHandler = echoExceptionFilter.Catch

}

func SetGlobalPrefix(e *echo.Echo) *echo.Group {
	return e.Group(config.GetConfiguration().Http.BasePath)
}

func EnableNotFound(e *echo.Echo, group *echo.Group) {
	e.Any("/*", customHTTP404RouterHandler)

	group.Any("/*", customHTTP404RouterHandler)

}

func EnableValidationRequest(e *echo.Echo) {
	//validator = validator.New()
	e.Validator = validate.NewGestGoValidator(validator.New())

}
func customV2HTTPErrorHandler(handlerErrors ...func(err error, c echo.Context)) {
	for _, handlerError := range handlerErrors {
		customV2HTTPErrorHandler(handlerError)
	}

}
func InternalServerErrorFilter(err error, c echo.Context) (code int, res any) {

	//he, ok := err.(*echo.HTTPError)
	// 400 status
	if he, ok := err.(*echo.HTTPError); ok {

		if he.Code == http.StatusBadRequest {
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
	//log.Print(err)
	//if _, ok := err.(*validator.InvalidValidationError); ok {
	//	fmt.Println(err)
	//	return
	//}

	errorRes := HttpError[string]{
		StatusCode: http.StatusInternalServerError,
		Message:    "Internal Server Error",
		Path:       c.Request().URL.Path,
		Timestamp:  time.Now().UnixMilli(),
	}
	return http.StatusInternalServerError, errorRes
}

func BadRequestErrorFilter(err error, c echo.Context) (code int, res any) {
	if he, ok := err.(*echo.HTTPError); ok {

		if he.Code == http.StatusBadRequest {
			errorBadRequest := BadRequestError[any]{
				HttpError: HttpError[any]{
					StatusCode: he.Code,
					Message:    "Invalid data format",
					Path:       c.Request().URL.Path,
					Timestamp:  time.Now().UnixMilli(),
				},
				Reasons: he.Message,
			}
			return http.StatusBadRequest, errorBadRequest
		}
	}
	return
}

func ValidateErrorFilter(err error, c echo.Context) (code int, res any) {
	if he, ok := err.(validator.ValidationErrors); ok {

		errorBadRequest := HttpError[any]{
			StatusCode: http.StatusBadRequest,
			Message:    he.Error(),
			Path:       c.Request().URL.Path,
			Timestamp:  time.Now().UnixMilli(),
		}
		return http.StatusBadRequest, errorBadRequest

	}
	return
}
func customHTTPErrorFilter(err error, c echo.Context) {

	//he, ok := err.(*echo.HTTPError)
	// 400 status
	if he, ok := err.(*echo.HTTPError); ok {

		if he.Code == http.StatusBadRequest {
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
	//log.Print(err)
	//if _, ok := err.(*validator.InvalidValidationError); ok {
	//	fmt.Println(err)
	//	return
	//}

	if he, ok := err.(validator.ValidationErrors); ok {
		error400 := HttpError[any]{
			StatusCode: http.StatusBadRequest,
			Message:    he.Error(),
			Path:       c.Request().URL.Path,
			Timestamp:  time.Now().UnixMilli(),
		}
		c.JSON(http.StatusBadRequest, error400)
		return
	}

	// 500 status
	errorRes := HttpError[string]{
		StatusCode: http.StatusInternalServerError,
		Message:    "Internal Server Error",
		Path:       c.Request().URL.Path,
		Timestamp:  time.Now().UnixMilli(),
	}
	c.JSON(http.StatusInternalServerError, errorRes)
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

type BadRequestError[T any] struct {
	HttpError[T]
	Reasons any `json:"reason,omitempty"`
}
type HttpError[T any] struct {
	StatusCode int    `json:"statusCode"`
	Message    T      `json:"message"`
	Path       string `json:"path"`
	Timestamp  int64  `json:"timestamp"`
}
