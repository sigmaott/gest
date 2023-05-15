package module

import (
	"fmt"
	errorGest "github.com/gestgo/gest/package/core/error"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"go.uber.org/fx"
	"net/http"
	validateMessage "payment/locales/validate-message"
	"strings"
	"time"
)

type I18nValidate struct {
	Ut       *ut.UniversalTranslator
	validate *validator.Validate
}
type I18nValidateParams struct {
	fx.In
	ut       *ut.UniversalTranslator
	validate *validator.Validate
}

func NewI18nValidate(validate *validator.Validate, Ut *ut.UniversalTranslator) *I18nValidate {

	lo.MapToSlice(validateMessage.ValidateMessage, func(key string, validateMessage map[string]string) bool {
		trans, found := Ut.GetTranslator(key)
		if !found {
			return found
		}
		for s, message := range validateMessage {
			validate.RegisterTranslation(s, trans, func(ut ut.Translator) error {
				return ut.Add(s, message, false)
			}, func(ut ut.Translator, fe validator.FieldError) string {
				//log.Print(fe.Field())
				t, _ := ut.T(s, fe.Field())
				return t
			})
		}
		return found
	})
	//enValidateMessage := validateMessage.ValidateMessage["en"]
	//for s, message := range enValidateMessage {
	//	validate.RegisterTranslation(s, trans, func(ut ut.Translator) error {
	//		return ut.Add(s, message, false)
	//	}, func(ut ut.Translator, fe validator.FieldError) string {
	//		//log.Print(fe.Field())
	//		t, _ := ut.T(s, fe.Field())
	//		return t
	//	})
	//}
	return &I18nValidate{
		Ut:       Ut,
		validate: validate,
	}
}

func (i *I18nValidate) ValidateErrorFilter(err error, c echo.Context) (code int, res any) {
	trans, found := i.Ut.GetTranslator(i.GetAcceptLanguage(c))
	if !found {
		trans = i.Ut.GetFallback()
	}
	if he, ok := err.(validator.ValidationErrors); ok {

		errorBadRequest := BadRequestError[any]{
			HttpError: errorGest.HttpError[any]{
				StatusCode: http.StatusBadRequest,
				Message:    "Bad Request",
				Path:       c.Request().URL.Path,
				Timestamp:  time.Now().UnixMilli(),
			},
			Errors: lo.MapToSlice(he.Translate(trans), func(key string, value string) string {
				return fmt.Sprintf("%s: %s", key[strings.Index(key, ".")+1:], value)
			}),
		}
		return http.StatusBadRequest, errorBadRequest

	}
	return
}
func (i *I18nValidate) GetAcceptLanguage(c echo.Context) string {
	language := c.Request().Header.Get("Accept-Language")
	if language == "" {
		return ""
	}
	languageSplit := strings.Split(language, "-")
	if len(languageSplit) == 2 {
		return languageSplit[0]
	}
	if len(languageSplit) == 1 {
		return languageSplit[0]
	}
	return ""
}
