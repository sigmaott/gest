package kafkafx

import (
	"context"
	"github.com/gestgo/gest/package/core/router"
	"go.uber.org/fx"
)

type Params struct {
	fx.In
	KafkaTopics []router.IRouter `group:"kafkaTopics"`
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
	return &KafkaSubscriber{}

}

type Result struct {
	fx.Out
	Topic router.IRouter `group:"kafkaTopics"`
}
