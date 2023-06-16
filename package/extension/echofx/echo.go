package echofx

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func ForRoot(port int) fx.Option {
	return fx.Module("echofx",

		fx.Provide(
			fx.Annotate(
				echo.New,
				fx.ResultTags(`name:"platformEcho"`)),
			fx.Annotate(
				func() int {
					return port
				},
				fx.ResultTags(`name:"platformEchoPort"`)),
			RegisterEchoHooks))
}
