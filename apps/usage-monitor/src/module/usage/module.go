package usage

import (
	"usage-monnitor/src/module/usage/controller"

	"github.com/gestgo/gest/package/extension/echofx"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("quota",
		fx.Provide(
			echofx.AsRoute(controller.NewAuthController),
		),
	)
}
