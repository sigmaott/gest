package http_generic

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockResponse struct {
	Message string `json:"message"`
}

type MockErrorResponse struct {
	Error string `json:"error"`
}

// Test the basic functionality with default JSON serialization/deserialization
func TestRequest_Success(t *testing.T) {
	mockResponse := MockResponse{Message: "Success"}
	mockResponseBody, _ := json.Marshal(mockResponse)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %v", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(mockResponseBody)
	}))
	defer server.Close()

	conf := JSONRequestConfig{
		Method:  http.MethodGet,
		URL:     server.URL,
		Headers: map[string]string{"Content-Type": "application/json"},
	}

	var result MockResponse
	result, err := Request[MockResponse](conf)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result.Message != "Success" {
		t.Errorf("Expected response message to be 'Success', got %v", result.Message)
	}
}

// Test the functionality with custom serialization/deserialization functions
func TestRequest_CustomSerialization(t *testing.T) {
	mockResponse := MockResponse{Message: "Custom Success"}
	mockResponseBody, _ := json.Marshal(mockResponse)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("Failed to read request body: %v", err)
		}
		defer r.Body.Close()

		var reqBody MockResponse
		json.Unmarshal(bodyBytes, &reqBody)

		if reqBody.Message != "Test Body" {
			t.Errorf("Expected request body message to be 'Test Body', got %v", reqBody.Message)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(mockResponseBody)
	}))
	defer server.Close()

	conf := JSONRequestConfig{
		Method: "POST",
		URL:    server.URL,
		Body:   MockResponse{Message: "Test Body"},
		Serialize: func(v any) ([]byte, error) {
			return json.Marshal(v)
		},
		Deserialize: func(body []byte, out any) error {
			return json.Unmarshal(body, out)
		},
	}

	var result MockResponse
	result, err := Request[MockResponse](conf)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result.Message != "Custom Success" {
		t.Errorf("Expected response message to be 'Custom Success', got %v", result.Message)
	}
}

// Test the case with custom client timeout
