package rabitmqfx

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("rabitmqfx", fx.Provide(RegisterRabbitmqHooks))
}
