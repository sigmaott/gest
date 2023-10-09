package kafkafx

import (
	"context"

	"github.com/sigmaott/gest/package/core/router"
	"go.uber.org/fx"
)

type Params struct {
	fx.In
	KafkaSubscriber *KafkaSubscriber `name:"platformKafka"`
	KafkaTopics     []any            `group:"kafkaTopics"`
}

func RegisterKafkaHooks(
	lifecycle fx.Lifecycle,
	params Params,
) *KafkaSubscriber {

	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go func() {
					router.InitRouter(params.KafkaTopics)
				}()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				return nil

			},
		})
	return params.KafkaSubscriber

}

type Result struct {
	fx.Out
	Topic router.IRouter `group:"kafkaTopics"`
}

func AsRoute(f any, annotation ...fx.Annotation) any {
	annotation = append(annotation, fx.As(new(any)),
		fx.ResultTags(`group:"kafkaTopics"`))
	return fx.Annotate(
		f,
		annotation...,
	)
}
