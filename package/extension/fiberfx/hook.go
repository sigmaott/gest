package fiberfx

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sigmaott/gest/package/core/router"
	"go.uber.org/fx"
)

type Params struct {
	fx.In
	PlatformFiber     *fiber.App `name:"platformFiber"`
	PlatformFiberPort int        `name:"platformFiberPort"`
	Routers           []any      `group:"fiberRouters"`
}

func RegisterFiberHooks(
	lifecycle fx.Lifecycle,
	params Params,
) *fiber.App {
	platformFiber := params.PlatformFiber

	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				router.InitRouter(params.Routers)
				go platformFiber.Listen(fmt.Sprintf(":%d", params.PlatformFiberPort))
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return platformFiber.Shutdown()
			},
		})
	return platformFiber
}

type Result struct {
	fx.Out
	Router router.IRouter `group:"fiberRouters"`
}

func AsRoute(f any, annotation ...fx.Annotation) any {
	annotation = append(annotation, fx.As(new(any)),
		fx.ResultTags(`group:"fiberRouters"`))
	return fx.Annotate(
		f,
		annotation...,
	)
}
