package rabitmqfx

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
)

type HandlerMessage func(message amqp.Delivery) error
type HandlerError func(err error) error
type RabbitmqSubscriber struct {
	Conn *amqp.Connection
}
type QueueConfig struct {
	Name                                string
	Consumer                            string
	AutoAck, Exclusive, NoLocal, noWait bool
	Args                                map[string]any
}

func (r *RabbitmqSubscriber) Subscribe(ctx context.Context, config QueueConfig, handler HandlerMessage, handlerErr HandlerError) {
	ch, err := r.Conn.Channel()
	if err != nil {
		handlerErr(err)
	}
	var a = &QueueConfig{}
	msgs, err := ch.Consume(
		config.Name,      // queue
		config.Consumer,  // consumer
		config.AutoAck,   // auto-ack
		config.Exclusive, // exclusive
		config.NoLocal,   // no-local
		config.noWait,    // no-wait
		a.Args,           // args
	)
	if err != nil {
		handlerErr(err)
	}
	go func() {

		for msg := range msgs {

			if err != nil {
				handlerErr(err)
			}
			err = handler(msg)

			if err != nil {
				handlerErr(err)
			}
		}
	}()

}
