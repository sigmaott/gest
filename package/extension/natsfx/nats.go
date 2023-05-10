package natsfx

import (
	"context"
	"github.com/gestgo/gest/package/core/router"
	"github.com/nats-io/nats.go"
	"go.uber.org/fx"
)

type Params struct {
	fx.In
	PlatformNats *nats.Conn
	Routers      []router.IRouter `group:"natsSubject"`
}

func RegisterNatsHooks(
	lifecycle fx.Lifecycle,
	params Params,
) *nats.Conn {
	platformNats := params.PlatformNats
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go func() {
					router.InitRouter(params.Routers)
				}()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				platformNats.Close()
				return nil

			},
		})
	return platformNats

}

type Result struct {
	fx.Out
	Router router.IRouter `group:"natsRouters"`
}
