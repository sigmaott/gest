package asynqfx

import (
	"context"
	"log"

	"github.com/hibiken/asynq"
	"github.com/sigmaott/gest/package/core/router"
	"go.uber.org/fx"
)

type Params struct {
	fx.In
	Server   *asynq.Server
	Jobs     []any `group:"asynqJobs"`
	ServeMux *asynq.ServeMux
}

type AsynqHook struct {
}

func registerAsynqHooks(
	lifecycle fx.Lifecycle,
	params Params,
) Result {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {

				router.InitRouter(params.Jobs)

				mux := params.ServeMux
				go func() error {
					if err := params.Server.Run(mux); err != nil {
						log.Fatalf("could not run server: %v", err)
					}
					return nil
				}()
				return nil
			},
			OnStop: func(context.Context) error {
				go params.Server.Stop()
				return nil
			},
		},
	)
	return Result{}
}

type Result struct {
}

func AsRoute(f any, annotation ...fx.Annotation) any {
	annotation = append(annotation, fx.As(new(any)),
		fx.ResultTags(`group:"asynqJobs"`))
	return fx.Annotate(
		f,
		annotation...,
	)
}
