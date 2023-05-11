package controller

import (
	"github.com/gestgo/gest/package/core/router"
	"github.com/gestgo/gest/package/extension/i18nfx"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"log"
	"net/http"
	"payment/src/module/payment/model"
	queryBuilder "payment/src/module/payment/query-builder"
	"payment/src/module/payment/repository"
)

type IUserController interface {
	FindAll()
}
type Params struct {
	fx.In
	Router      *echo.Group
	Logger      *zap.SugaredLogger
	I18nService i18nfx.II18nService
}
type Controller struct {
	router      *echo.Group
	logger      *zap.SugaredLogger
	i18nService i18nfx.II18nService
	repository  repository.IPaymentRepository
}

type Result struct {
	fx.Out
	Controller router.IRouter `group:"echoRouters"`
}

func NewController(params Params) IUserController {
	return &Controller{
		router:      params.Router,
		logger:      params.Logger,
		i18nService: params.I18nService,
	}
}

func NewRouter(params Params) Result {
	c := NewController(params)
	return Result{Controller: router.NewBaseRouter[IUserController](c)}

}

func (b *Controller) FindAll() {
	b.router.GET("/users", func(c echo.Context) error {

		//message, err := b.i18nService.T("en", locales.CARDINAL_TEST)
		result, sort, err := queryBuilder.MongoParserQuery[model.Payment](c.Request().URL.Query())
		log.Print(result, sort, err)
		//b.logger.Info()
		//b.repository.FindAll()
		return c.String(http.StatusOK, "ok")
	})

}
