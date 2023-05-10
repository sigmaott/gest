package module

import (
	"github.com/gestgo/gest/package/extension/echofx/exception"
	"github.com/gestgo/gest/package/technique/validate"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"payment/config"
	"payment/docs"
	"payment/src/echoSwagger"
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

func EnableErrorHandler(e *echo.Echo, exception exceptions.IEchoCustomException) {
	e.HTTPErrorHandler = exception.ErrorHandler
}

func SetGlobalPrefix(e *echo.Echo) *echo.Group {
	return e.Group(config.GetConfiguration().Http.BasePath)
}
