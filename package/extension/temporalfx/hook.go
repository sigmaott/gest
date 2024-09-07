package temporalfx

import (
	"context"

	"github.com/sigmaott/gest/package/core/router"
	"go.temporal.io/sdk/client"
	"go.uber.org/fx"
)

type Params struct {
	fx.In
	TemporalClient  client.Client `name:"temporalClient"`
	WorkerFactories []any         `group:"temporalWorkers"`
}

func RegisterTemporalHooks(
	lifecycle fx.Lifecycle,
	params Params,
) *Result {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go func() {
					router.InitRouter(params.WorkerFactories)
				}()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				return nil

			},
		})
	return &Result{}

}

type Result struct {
}

func AsRoute(f any, annotation ...fx.Annotation) any {
	annotation = append(annotation, fx.As(new(any)),
		fx.ResultTags(`group:"temporalWorkers"`))
	return fx.Annotate(
		f,
		annotation...,
	)
}
