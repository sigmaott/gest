package asynqfx

import (
	"context"
	"github.com/gestgo/gest/package/core/router"
	"github.com/hibiken/asynq"
	"go.uber.org/fx"
	"log"
)

type Params struct {
	fx.In
	Server   *asynq.Server
	Jobs     []any `group:"asynqJobs"`
	ServeMux *asynq.ServeMux
}

type AsynqHook struct {
}

func RegisterAsynqHooks(
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
	return Result{
		AsynqHook: &AsynqHook{},
	}
}

type Result struct {
	fx.Out
	AsynqHook *AsynqHook
}
