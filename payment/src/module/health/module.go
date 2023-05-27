package health

import (
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("health",
		fx.Provide(
			NewHealthRouter,
			NewHeathCheckService,
		),
	)
}
