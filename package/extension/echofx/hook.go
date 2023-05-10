package echofx

import (
	"context"
	"fmt"
	"github.com/gestgo/gest/package/core/router"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type Params struct {
	fx.In
	PlatformEcho     *echo.Echo       `name:"platformEcho"`
	PlatformEchoPort int              `name:"platformEchoPort"`
	Routers          []router.IRouter `group:"echoRouters"`
}

func RegisterEchoHooks(
	lifecycle fx.Lifecycle,
	params Params,
) *echo.Echo {
	platformEcho := params.PlatformEcho

	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				router.InitRouter(params.Routers)
				go platformEcho.Start(fmt.Sprintf(":%d", params.PlatformEchoPort))
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return platformEcho.Shutdown(ctx)

			},
		})
	return platformEcho

}

type Result struct {
	fx.Out
	Router router.IRouter `group:"echoRouters"`
}
