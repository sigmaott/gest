package kafkafx

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("kafkafx",
		fx.Provide(
			fx.Annotate(
				func() *KafkaSubscriber {
					return &KafkaSubscriber{}
				},
				fx.ResultTags(`name:"platformKafka"`))),

		fx.Provide(RegisterKafkaHooks))
}
