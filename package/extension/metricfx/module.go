package metricfx

import (
	"go.uber.org/fx"
)

func ForRoot() fx.Option {
	return fx.Module("metricfx",
		fx.Provide(NewPrometheusMetrics),
	)
}
