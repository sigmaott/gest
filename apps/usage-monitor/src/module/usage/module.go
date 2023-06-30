package usage

import (
	"usage-monnitor/src/module/usage/controller"
	"usage-monnitor/src/module/usage/repository"
	"usage-monnitor/src/module/usage/service"

	"github.com/gestgo/gest/package/extension/grpcfx"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("quota",
		fx.Provide(
			grpcfx.AsRoute(controller.NewQuotaGrpcController()),
		),
		fx.Provide(
			service.NewSSAIUsageMonitorService,
		),
		fx.Provide(
			repository.NewIQuotaRepository,
			repository.NewIQuotaRepository,
		),
	)
}
