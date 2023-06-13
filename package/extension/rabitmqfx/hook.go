package rabitmqfx

import (
	"context"
	"github.com/gestgo/gest/package/core/router"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/fx"
)

type Params struct {
	fx.In
	PlatformRabbitmqMQ *amqp.Connection `name:"platformRabbitMQ"`
	RabbitQueues       []any            `group:"rabbitQueues"`
}

func RegisterRabbitmqHooks(
	lifecycle fx.Lifecycle,
	params Params,
) *amqp.Connection {

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
	return platformRabbitmqMQ

}

type ParamsNewRabbitmqSubscriber struct {
	fx.In
	PlatformRabbitmqMQ *amqp.Connection `name:"platformRabbitMQ"`
}

func NewRabbitmqSubscriber(params ParamsNewRabbitmqSubscriber) *RabbitmqSubscriber {

	return &RabbitmqSubscriber{
		Conn: params.PlatformRabbitmqMQ,
	}

}
