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
			func() client.Client {
				return temporalClient
			},
			fx.Annotate(
				func() *Result {
					return &Result{}
				},
				fx.ResultTags(`group:"temporalWorkers"`),
			),
		),
		fx.Provide(
			worker_factory.NewWorkerFactory,
		),
		// Register lifecycle hooks
		fx.Provide(
			RegisterTemporalHooks,
		),
	)
}
