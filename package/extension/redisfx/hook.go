package redisfx

import (
	"context"
	"github.com/gestgo/gest/package/core/router"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
)

type Params struct {
	fx.In
	PlatformRedis redis.UniversalClient
	Channels      []router.IRouter `group:"redisChannels"`
}

func RegisterRedisHooks(
	lifecycle fx.Lifecycle,
	params Params,
) *RedisSubscriber {
	platformRedis := params.PlatformRedis
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				go func() {
					router.InitRouter(params.Channels)
				}()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				return platformRedis.Close()

			},
		})
	return &RedisSubscriber{
		Client: platformRedis,
	}

}

type Result struct {
	fx.Out
	Channel router.IRouter `group:"redisChannels"`
}
