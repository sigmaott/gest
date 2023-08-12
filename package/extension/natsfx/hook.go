package natsfx

import (
	"go.uber.org/fx"
)

func ForRoot() fx.Option {
	return fx.Module("natsfx", fx.Provide(RegisterNatsHooks))
}
