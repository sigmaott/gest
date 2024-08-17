package http_generic

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type SerializationFunc func(any) ([]byte, error)
type DeserializationFunc func(body []byte, out any) error

type JSONRequestConfig struct {
	Method      string
	URL         string
	Headers     map[string]string
	Body        any
	Client      *http.Client
	Timeout     *time.Duration
	Serialize   SerializationFunc
	Deserialize DeserializationFunc
}

func Request[R any](conf JSONRequestConfig) (R, error) {
	var result R

	// Set default client if none provided
	if conf.Client == nil {
		conf.Client = &http.Client{}
		if conf.Timeout != nil {
			conf.Client.Timeout = *conf.Timeout
		}
	}

	// Serialize the request body if provided
	var bodyBytes []byte
	var err error
	if conf.Body != nil {
		if conf.Serialize != nil {
			bodyBytes, err = conf.Serialize(conf.Body)
		} else {
			bodyBytes, err = json.Marshal(conf.Body)
		}
		if err != nil {
			return result, err
		}
	}

	// Create HTTP request
	req, err := http.NewRequest(conf.Method, conf.URL, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return result, err
	}

	// Set headers
	// req.Header.Set("Content-Type", "application/json")
	for header, value := range conf.Headers {

		req.Header.Set(header, value)
	}

	// Execute the request

	resp, err := conf.Client.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	// Deserialize the response body
	if conf.Deserialize != nil {
		err = conf.Deserialize(respBody, &result)
	} else {
		err = json.Unmarshal(respBody, &result)
	}
	return result, err
}
