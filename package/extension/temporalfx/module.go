package temporalfx

import (
	"context"

	worker_factory "github.com/sigmaott/gest/package/extension/temporalfx/worker-factory"
	"go.temporal.io/sdk/client"
	"go.uber.org/fx"
)

// ForRoot returns an fx.Option that sets up the Temporal module.
func ForRoot(ctx context.Context, temporalClient client.Client) fx.Option {
	return fx.Module("temporalfx",
		// Provide the Temporal client with a name
		fx.Provide(
			fx.Annotate(
				func() client.Client {
					return temporalClient
				},
				fx.ResultTags(`name:"temporalClient"`),
			),
		),
		fx.Provide(
			fx.Annotate(
				worker_factory.NewWorkerFactory,
				fx.ParamTags(`name:"temporalClient"`),
			),
		),
		// Register lifecycle hooks
		fx.Provide(
			RegisterTemporalHooks,
		),
	)
}
