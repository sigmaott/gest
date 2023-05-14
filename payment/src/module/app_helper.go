package module

import (
	"fmt"
	errorGest "github.com/gestgo/gest/package/core/error"
	"github.com/gestgo/gest/package/technique/validate"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	viTranslations "github.com/go-playground/validator/v10/translations/vi"
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

func EnableErrorHandler(e *echo.Echo, i18nValidate *I18nValidate) {
	echoExceptionFilter := errorGest.NewEchoExceptionFilter(BadRequestErrorFilter, i18nValidate.ValidateErrorFilter, InternalServerErrorFilter)
	e.HTTPErrorHandler = echoExceptionFilter.CatchError

}

func SetGlobalPrefix(e *echo.Echo) *echo.Group {
	return e.Group(config.GetConfiguration().Http.BasePath)
}

func EnableNotFound(e *echo.Echo, group *echo.Group) {
	e.Any("/*", customHTTP404RouterHandler)

	//group.Any("/*", customHTTP404RouterHandler)

}

func EnableValidationRequest(e *echo.Echo, v *validator.Validate) {
	e.Validator = validate.NewGestGoValidator(v)

}

func InternalServerErrorFilter(err error, c echo.Context) (code int, res any) {

	//he, ok := err.(*echo.HTTPError)
	// 400 status
	//if he, ok := err.(*echo.HTTPError); ok {
	//
	//	if he.Code == http.StatusBadRequest {
	//		error400 := errorGest.HttpError[any]{
	//			StatusCode: he.Code,
	//			Message:    he.Message,
	//			Path:       c.Request().URL.Path,
	//			Timestamp:  time.Now().UnixMilli(),
	//		}
	//		c.JSON(http.StatusBadRequest, error400)
	//		return
	//	}
	//}
	log.Print(err)
	errorRes := errorGest.HttpError[string]{
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
				HttpError: errorGest.HttpError[any]{
					StatusCode: he.Code,
					Message:    "Bad Request",
					Path:       c.Request().URL.Path,
					Timestamp:  time.Now().UnixMilli(),
				},
				Errors: he.Message,
			}
			return http.StatusBadRequest, errorBadRequest
		}
	}
	return
}

func customHTTP404RouterHandler(c echo.Context) error {
	code := http.StatusNotFound
	errorRes := errorGest.HttpError[string]{
		StatusCode: code,
		Message:    fmt.Sprintf("Cannot %s %s", c.Request().Method, c.Request().URL.Path),
		Path:       c.Request().URL.Path,
		Timestamp:  time.Now().UnixMilli(),
	}
	return c.JSON(code, errorRes)
}

func EnableLogRouter(e *echo.Echo, logger *zap.SugaredLogger) {
	//data, _ := json.MarshalIndent(e.Routes(), "", "  ")
	//
	////e.Routers()
	//logger.Infof("%+v", string(data))
	logger.Infof("%+v", "******************* Router ***************")
	for _, route := range e.Routes() {
		logger.Infof(" %+v %+v", route.Method, route.Path)
	}

}

func RegisterValidateTranslations(validate *validator.Validate, Ut *ut.UniversalTranslator) {
	enTrans, _ := Ut.GetTranslator("en")
	err := enTranslations.RegisterDefaultTranslations(validate, enTrans)
	if err != nil {

		log.Print(err)
		return
	}
	viTrans, _ := Ut.GetTranslator("vi")
	err = viTranslations.RegisterDefaultTranslations(validate, viTrans)
	if err != nil {
		log.Print(err)
		return
	}
}
