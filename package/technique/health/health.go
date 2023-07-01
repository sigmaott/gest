package healthcheck

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"sync"
	"time"
)

type Response struct {
	Status string            `json:"status,omitempty"`
	Errors map[string]string `json:"errors,omitempty"`
}

type Health struct {
	Checkers  map[string]Checker
	Observers map[string]Checker
	Timeout   time.Duration
}

// Checker checks the status of the dependency and returns error.
// In case the dependency is working as expected, return nil.
type Checker interface {
	Check(ctx context.Context) error
}

// CheckerFunc is a convenience type to create functions that implement the Checker interface.go.
type CheckerFunc func(ctx context.Context) error

// Check Implements the Checker interface.go to allow for any func() error method
// to be passed as a Checker
func (c CheckerFunc) Check(ctx context.Context) error {
	return c(ctx)
}

// Handler returns an http.Handler
func Handler(opts ...Option) http.Handler {
	h := &Health{
		Checkers:  make(map[string]Checker),
		Observers: make(map[string]Checker),
		Timeout:   30 * time.Second,
	}
	for _, opt := range opts {
		opt(h)
	}
	return h
}

// HandlerFunc returns an http.HandlerFunc to mount the API implementation at a specific route
func HandlerFunc(opts ...Option) http.HandlerFunc {
	return Handler(opts...).ServeHTTP
}

// Option adds optional parameter for the HealthcheckHandlerFunc
type Option func(*Health)

// WithChecker adds a status checker that needs to be added as part of healthcheck. i.e database, cache or any external dependency
func WithChecker(name string, s Checker) Option {
	return func(h *Health) {
		h.Checkers[name] = &timeoutChecker{s}
	}
}

// WithObserver adds a status checker but it does not fail the entire status.
func WithObserver(name string, s Checker) Option {
	return func(h *Health) {
		h.Observers[name] = &timeoutChecker{s}
	}
}

// WithTimeout configures the global Timeout for all individual Checkers.
func WithTimeout(timeout time.Duration) Option {
	return func(h *Health) {
		h.Timeout = timeout
	}
}

func (h *Health) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	nCheckers := len(h.Checkers) + len(h.Observers)

	code := http.StatusOK
	errorMsgs := make(map[string]string, nCheckers)

	ctx, cancel := context.Background(), func() {}
	if h.Timeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, h.Timeout)
	}
	defer cancel()

	var mutex sync.Mutex
	var wg sync.WaitGroup
	wg.Add(nCheckers)

	for key, checker := range h.Checkers {
		go func(key string, checker Checker) {
			if err := checker.Check(ctx); err != nil {
				mutex.Lock()
				errorMsgs[key] = err.Error()
				code = http.StatusServiceUnavailable
				mutex.Unlock()
			}
			wg.Done()
		}(key, checker)
	}
	for key, observer := range h.Observers {
		go func(key string, observer Checker) {
			if err := observer.Check(ctx); err != nil {
				mutex.Lock()
				errorMsgs[key] = err.Error()
				mutex.Unlock()
			}
			wg.Done()
		}(key, observer)
	}

	wg.Wait()

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(Response{
		Status: http.StatusText(code),
		Errors: errorMsgs,
	})
}

type timeoutChecker struct {
	checker Checker
}

func (t *timeoutChecker) Check(ctx context.Context) error {
	checkerChan := make(chan error)
	go func() {
		checkerChan <- t.checker.Check(ctx)
	}()
	select {
	case err := <-checkerChan:
		return err
	case <-ctx.Done():
		return errors.New("max check time exceeded")
	}
}

func HandlerHeathCheck(opts ...Option) *Response {
	h := &Health{
		Checkers:  make(map[string]Checker),
		Observers: make(map[string]Checker),
		Timeout:   30 * time.Second,
	}
	for _, opt := range opts {
		opt(h)
	}
	nCheckers := len(h.Checkers) + len(h.Observers)

	code := http.StatusOK
	errorMsgs := make(map[string]string, nCheckers)

	ctx, cancel := context.Background(), func() {}
	if h.Timeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, h.Timeout)
	}
	defer cancel()

	var mutex sync.Mutex
	var wg sync.WaitGroup
	wg.Add(nCheckers)

	for key, checker := range h.Checkers {
		go func(key string, checker Checker) {
			if err := checker.Check(ctx); err != nil {
				mutex.Lock()
				errorMsgs[key] = err.Error()
				code = http.StatusServiceUnavailable
				mutex.Unlock()
			}
			wg.Done()
		}(key, checker)
	}
	for key, observer := range h.Observers {
		go func(key string, observer Checker) {
			if err := observer.Check(ctx); err != nil {
				mutex.Lock()
				errorMsgs[key] = err.Error()
				mutex.Unlock()
			}
			wg.Done()
		}(key, observer)
	}

	wg.Wait()

	return &Response{
		Status: http.StatusText(code),
		Errors: errorMsgs,
	}
}
