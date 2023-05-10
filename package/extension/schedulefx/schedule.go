package schedulefx

import (
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("schedulefx", fx.Provide(RegisterScheduleHooks))
}
