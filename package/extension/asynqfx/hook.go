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
	Server *asynq.Server
	Jobs   []router.IRouter `group:"asyncJobs"`
}

func RegisterAsynqHooks(
	lifecycle fx.Lifecycle,
	params Params,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {

				router.InitRouter(params.Jobs)
				mux := asynq.NewServeMux()
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
}
