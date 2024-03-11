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
) *Result {

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
	return &Result{}

}

type Result struct {
}

func AsRoute(f any, annotation ...fx.Annotation) any {
	annotation = append(annotation, fx.As(new(any)),
		fx.ResultTags(`group:"kafkaTopics"`))
	return fx.Annotate(
		f,
		annotation...,
	)
}
