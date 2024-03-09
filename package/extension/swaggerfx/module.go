package swagger

import (
	"github.com/go-swagno/swagno"
	"go.uber.org/fx"
)

func ForRoot(conf swagno.Config) fx.Option {

	return fx.Module("swaggerfx", fx.Provide(
		func() *swagno.Swagger {
			return swagno.New(conf)
		},
	))

}
