package natsfx

import (
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("natsfx", fx.Provide(RegisterNatsHooks))
}
