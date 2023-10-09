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
func AsRoute(f any, annotation ...fx.Annotation) any {
	annotation = append(annotation, fx.As(new(any)),
		fx.ResultTags(`group:"natsSubject"`))
	return fx.Annotate(
		f,
		annotation...,
	)
}
