package health

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
)

type IHealthController interface {
	Health()
}
type Params struct {
	fx.In
	Router  *echo.Group
	Logger  *zap.SugaredLogger
	Service IHeathCheckService
}
type Controller struct {
	router  *echo.Group
	logger  *zap.SugaredLogger
	service IHeathCheckService
}

func (h *Controller) Health() {
	h.router.GET("/health", func(c echo.Context) error {
		res, err := h.service.HeathCheck()
		if err != nil {
			return c.JSON(http.StatusServiceUnavailable, res)
		}
		return c.JSON(http.StatusOK, res)
	})
}

type Result struct {
	fx.Out
	Controller any `group:"echoRouters"`
}

func NewHealthController(params Params) Result {
	return Result{
		Controller: &Controller{
			router:  params.Router,
			logger:  params.Logger,
			service: params.Service,
		},
	}
}

//func NewHealthRouter(params Params) Result {
//	c := NewHealthController(params)
//	return Result{Controller: router.NewBaseRouter[IHealthController](c)}
//}
