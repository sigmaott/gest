package redisfx

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type HandlerMessage func(message *redis.Message) error
type HandlerError func(err error) error
type RedisSubscriber struct {
	Client redis.UniversalClient
}

func (c *RedisSubscriber) Subscribe(ctx context.Context, handler HandlerMessage, handlerErr HandlerError, channels ...string) {
	subscriber := c.Client.Subscribe(ctx, channels...)
	go c.handler(subscriber, ctx, handler, handlerErr)

}
func (c *RedisSubscriber) handler(subscriber *redis.PubSub, ctx context.Context, handler HandlerMessage, handlerErr HandlerError) {
	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			handlerErr(err)
		}
		err = handler(msg)

		if err != nil {
			handlerErr(err)
		}
	}
}
