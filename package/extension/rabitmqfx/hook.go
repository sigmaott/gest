package rabitmqfx

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sigmaott/gest/package/core/router"
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
func AsRoute(f any, annotation ...fx.Annotation) any {
	annotation = append(annotation, fx.As(new(any)),
		fx.ResultTags(`group:"rabbitQueues"`))
	return fx.Annotate(
		f,
		annotation...,
	)
}
