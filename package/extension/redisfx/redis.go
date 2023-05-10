package redisfx

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("redisfx", fx.Provide(RegisterRedisHooks))
}
