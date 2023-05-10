package asynqfx

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("asynqfx", fx.Provide(RegisterAsynqHooks))
}
