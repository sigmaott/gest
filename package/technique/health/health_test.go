package healthcheck

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestNewHandlerFunc(t *testing.T) {
	//type args struct {
	//	opts []Option
	//}
	tests := []struct {
		name       string
		args       []Option
		statusCode int
		response   Response
	}{
		{
			name:       "returns 200 status if no errors",
			statusCode: http.StatusOK,
			response: Response{
				Status: http.StatusText(http.StatusOK),
			},
		},
		{
			name:       "returns 503 status if errors",
			statusCode: http.StatusServiceUnavailable,
			args: []Option{
				WithChecker("database", CheckerFunc(func(ctx context.Context) error {
					return fmt.Errorf("connection to db timed out")
				})),
				WithChecker("testService", CheckerFunc(func(ctx context.Context) error {
					return fmt.Errorf("connection refused")
				})),
			},
			response: Response{
				Status: http.StatusText(http.StatusServiceUnavailable),
				Errors: map[string]string{
					"database":    "connection to db timed out",
					"testService": "connection refused",
				},
			},
		},
		{
			name:       "returns 503 status if Checkers Timeout",
			statusCode: http.StatusServiceUnavailable,
			args: []Option{
				WithTimeout(1 * time.Millisecond),
				WithChecker("database", CheckerFunc(func(ctx context.Context) error {
					time.Sleep(10 * time.Millisecond)
					return nil
				})),
			},
			response: Response{
				Status: http.StatusText(http.StatusServiceUnavailable),
				Errors: map[string]string{
					"database": "max check time exceeded",
				},
			},
		},
		{
			name:       "returns 200 status if errors are observable",
			statusCode: http.StatusOK,
			args: []Option{
				WithObserver("observableService", CheckerFunc(func(ctx context.Context) error {
					return fmt.Errorf("i fail but it is okay")
				})),
			},
			response: Response{
				Status: http.StatusText(http.StatusOK),
				Errors: map[string]string{
					"observableService": "i fail but it is okay",
				},
			},
		},
		{
			name:       "returns 503 status if errors with observable fails",
			statusCode: http.StatusServiceUnavailable,
			args: []Option{
				WithObserver("database", CheckerFunc(func(ctx context.Context) error {
					return fmt.Errorf("connection to db timed out")
				})),
				WithChecker("testService", CheckerFunc(func(ctx context.Context) error {
					return fmt.Errorf("connection refused")
				})),
			},
			response: Response{
				Status: http.StatusText(http.StatusServiceUnavailable),
				Errors: map[string]string{
					"database":    "connection to db timed out",
					"testService": "connection refused",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := httptest.NewRecorder()
			req, err := http.NewRequest("GET", "http://localhost/health", nil)
			if err != nil {
				t.Errorf("Failed to create request.")
			}
			HandlerFunc(tt.args...)(res, req)
			if res.Code != tt.statusCode {
				t.Errorf("expected code %d, got %d", tt.statusCode, res.Code)
			}
			var respBody Response
			if err := json.NewDecoder(res.Body).Decode(&respBody); err != nil {
				t.Fatal("failed to parse the body")
			}
			if !reflect.DeepEqual(respBody, tt.response) {
				t.Errorf("NewHandlerFunc() = %v, want %v", respBody, tt.response)
			}
		})
	}
}
