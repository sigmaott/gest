package payment

import (
	//"github.com/gestgo/gest/package/core/repository"
	"go.uber.org/fx"
	"payment/src/module/payment/controller"
	"payment/src/module/payment/repository"
	"payment/src/module/payment/service"
)

func Module() fx.Option {
	return fx.Module("payment",
		fx.Provide(
			controller.NewRouter,
			service.NewUserService,
			repository.NewPaymentRepository,
		),
	)
}

//fx.ResultTags(`group:"controllers"`)
