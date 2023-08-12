package natsfx

import (
	"context"

	"github.com/nats-io/nats.go"
	"github.com/sigmaott/gest/package/core/router"
	"go.uber.org/fx"
)

type Params struct {
	fx.In
	PlatformNats *nats.Conn
	Routers      []any `group:"natsSubject"`
}

func RegisterNatsHooks(
	lifecycle fx.Lifecycle,
	params Params,
) *NatsHook {
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
	return &NatsHook{
		platformNats,
	}

}

type NatsHook struct {
	PlatformNats *nats.Conn
}

func AsRoute(f any, annotation ...fx.Annotation) any {
	annotation = append(annotation, fx.As(new(any)),
		fx.ResultTags(`group:"natsSubject"`))
	return fx.Annotate(
		f,
		annotation...,
	)
}
