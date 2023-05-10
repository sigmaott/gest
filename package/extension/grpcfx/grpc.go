package grpcfx

import (
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("grpcfx", fx.Provide(RegisterGRPCHooks))
}
