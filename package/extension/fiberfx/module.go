package fiberfx

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func ForRoot(port int, conf fiber.Config) fx.Option {
	return fx.Module("fiberfx",

		fx.Provide(
			fx.Annotate(
				func() *fiber.App {
					return fiber.New(conf)
				}, // Create a new Fiber instance
				fx.ResultTags(`name:"platformFiber"`)), // Provide it with the name "platformFiber"
			fx.Annotate(
				func() int {
					return port
				},
				fx.ResultTags(`name:"platformFiberPort"`)), // Provide the port number with the name "platformFiberPort"
			RegisterFiberHooks)) // Register the Fiber lifecycle hooks
}
