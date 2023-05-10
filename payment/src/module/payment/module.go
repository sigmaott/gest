package payment

import (
	"go.uber.org/fx"
	"payment/src/module/payment/controller"
	"payment/src/module/payment/service"
)

func Module() fx.Option {
	return fx.Module("payment",
		fx.Provide(
			controller.NewRouter,
			service.NewUserService,
		),
	)
}

//fx.ResultTags(`group:"controllers"`)
