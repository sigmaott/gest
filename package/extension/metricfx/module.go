package metricfx

import "go.uber.org/fx"

// Module exports Prometheus metrics as a dependency.
var Module = fx.Options(
	fx.Provide(NewPrometheusMetrics),
)
