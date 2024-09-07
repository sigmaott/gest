package worker_factory

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

// IWorkerFactory defines an interface for initializing workers.
type IWorkerFactory interface {
	InitializeWorker(taskQueue string, opts worker.Options, handlers ...WorkerHandler)
}

// WorkerFactory defines a struct to create Temporal workers
type WorkerFactory struct {
	Client client.Client
}

// NewWorkerFactory creates a new WorkerFactory
func NewWorkerFactory(c client.Client) IWorkerFactory {
	return &WorkerFactory{
		Client: c,
	}
}

// WorkerHandler defines a function type for registering workflows and activities
type WorkerHandler func(worker worker.Worker)

// InitializeWorker creates a new worker with the provided options and registers handlers
func (wf *WorkerFactory) InitializeWorker(taskQueue string, opts worker.Options, handlers ...WorkerHandler) {
	// Create a new worker
	w := worker.New(wf.Client, taskQueue, opts)

	// Apply each handler to the worker
	for _, handler := range handlers {
		handler(w)
	}

	// Start the worker and return any errors encountered
	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln("Unable to start worker", err)
	}

}
