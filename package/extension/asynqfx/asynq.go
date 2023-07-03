package asynqfx

import (
	"github.com/hibiken/asynq"
	"go.uber.org/fx"
)

//
//func Module() fx.Option {
//	return fx.Module("asynqfx", fx.Provide(registerAsynqHooks))
//}

type Param struct {
	Server   *asynq.Server
	ServeMux *asynq.ServeMux
}

func ForRoot(param Param) fx.Option {
	return fx.Module("asynqfx",
		fx.Provide(func() *asynq.Server {
			return param.Server
		}),
		fx.Provide(func() *asynq.ServeMux {
			return param.ServeMux
		}),
		fx.Provide(registerAsynqHooks))
}
