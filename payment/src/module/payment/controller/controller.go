package controller

import (
	"github.com/gestgo/gest/package/core/router"
	"github.com/gestgo/gest/package/extension/echofx/parser"
	"github.com/gestgo/gest/package/extension/i18nfx"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
	"payment/locales"
	"payment/src/module/payment/dto"
)

type IUserController interface {
	Create()
	FindOne()
	FindAll()
	Update()
	Delete()
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

func (b *Controller) Create() {

	b.router.POST("/users", func(c echo.Context) error {
		body := c.Get("body").(*dto.CreateUser)
		return c.JSON(http.StatusOK, body)
	}, parser.NewBodyParser[dto.CreateUser]("body", true).Parser)

}

func (b *Controller) FindAll() {
	b.router.GET("/users", func(c echo.Context) error {

		message, err := b.i18nService.T("en", locales.CARDINAL_TEST)
		b.logger.Info(err)
		return c.String(http.StatusOK, message)
	})

}

func (b *Controller) FindOne() {

	b.router.GET("/users/:id", func(c echo.Context) error {

		u := c.Get("param").(*dto.GetUserById)
		return c.JSON(http.StatusOK, u)
	}, parser.NewParamsParser[dto.GetUserById]("param", true).Parser)

}
func (b *Controller) Update() {

	b.router.PUT("/users/:id", func(c echo.Context) error {

		u := c.Get("request").(*dto.UpdateUser)
		return c.JSON(http.StatusOK, u)
	}, parser.NewRequestParser[dto.UpdateUser]("request", true).Parser)
}

func (b *Controller) Delete() {
	b.router.DELETE("/users/:id", func(c echo.Context) error {
		u := c.Get("request").(*dto.DeleteUserById)
		return c.JSON(http.StatusOK, u)
	}, parser.NewRequestParser[dto.DeleteUserById]("request", true).Parser)
}
