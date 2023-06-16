package usage

import (
	"github.com/gestgo/gest/package/extension/grpcfx"
	"go.uber.org/fx"
	"quota/src/module/quota/controller"
	"quota/src/module/quota/repository"
	"quota/src/module/quota/service"
)

func Module() fx.Option {
	return fx.Module("quota",
		fx.Provide(
			grpcfx.AsRoute(controller.NewQuotaGrpcController),
		),
		fx.Provide(
			service.NewQuotaService,
		),
		fx.Provide(
			repository.NewBaseQuotaRepository,
			repository.NewAppQuotaRepository,
		),
	)
}
