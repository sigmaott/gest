package schedulefx

import (
	"context"
	"github.com/gestgo/gest/package/core/router"
	"github.com/go-co-op/gocron"
	"go.uber.org/fx"
)

type Params struct {
	fx.In
	Platform *gocron.Scheduler `name:"platformGoCron"`
	CronJobs []router.IRouter  `group:"cronJobs"`
}

func RegisterScheduleHooks(
	lifecycle fx.Lifecycle,
	params Params,
) Result {
	platform := params.Platform
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go func() {
					router.InitRouter(params.CronJobs)
					platform.StartBlocking()
				}()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				platform.Clear()
				return nil

			},
		})
	return Result{
		Platform: platform,
	}

}

type Result struct {
	fx.Out
	Platform *gocron.Scheduler
}
