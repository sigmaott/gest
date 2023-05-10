package rabitmqfx

import (
	"context"
	"github.com/gestgo/gest/package/core/router"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/fx"
)

type Params struct {
	fx.In
	PlatformRabbitmqMQ *amqp.Connection
	RabbitQueues       []router.IRouter `group:"rabbitQueues"`
}

func RegisterRabbitmqHooks(
	lifecycle fx.Lifecycle,
	params Params,
) *RabbitmqSubscriber {

	platformRabbitmqMQ := params.PlatformRabbitmqMQ
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go func() {
					router.InitRouter(params.RabbitQueues)
				}()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				return platformRabbitmqMQ.Close()

			},
		})
	return &RabbitmqSubscriber{
		Conn: platformRabbitmqMQ,
	}

}

type Result struct {
	fx.Out
	Channel router.IRouter `group:"redisChannels"`
}
