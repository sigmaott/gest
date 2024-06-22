package metricfx

import (
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/fx"
)

func ForRoot(registry *prometheus.Registry) fx.Option {
	return fx.Module("metricfx",
		fx.Provide(NewPrometheusMetrics),
	)
}
