package kafkafx

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("kafkafx", fx.Provide(RegisterKafkaHooks))
}
