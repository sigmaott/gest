package controller

import (
	"log"
	"net/http"

	"github.com/gestgo/gest/package/extension/i18nfx"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type IAuthController interface {
	Auth()
}

type Controller struct {
	router      *echo.Group
	logger      *zap.SugaredLogger
	i18nService i18nfx.II18nService
}

func (h *Controller) Auth() {
	h.router.GET("/auth", func(c echo.Context) error {
		message, err := h.i18nService.T("vi", "cardinal_test")
		log.Print(err)
		return c.String(http.StatusOK, message)
	})
}

func NewAuthController(router *echo.Group, i18nService i18nfx.II18nService,
	logger *zap.SugaredLogger) IAuthController {
	return &Controller{
		router:      router,
		logger:      logger,
		i18nService: i18nService,
	}
}
