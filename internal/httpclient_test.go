package internal

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {

	tests := []struct {
		name            string
		options         []func(*Option)
		expectedOptions *Option
		expectClient    bool
	}{
		{
			name:    "Should create client with default options",
			options: []func(*Option){},
			expectedOptions: &Option{
				blocking:      false,
				limit:         defaultLimit,
				limitInterval: defaultInterval,
			},
			expectClient: true,
		},
		{
			name: "Should create client with custom options",
			options: []func(*Option){
				WithBlocking(),
				WithLimit(10),
				WithLimitInterval(time.Second),
			},
			expectedOptions: &Option{
				blocking:      true,
				limit:         10,
				limitInterval: time.Second,
			},
			expectClient: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hc := NewClient(tt.options...)
			if tt.expectClient {
				assert.NotNil(t, hc)
				assert.NotNil(t, hc.client)
				assert.NotNil(t, hc.rateLimiter)
			} else {
				assert.Nil(t, hc)
			}
			assert.EqualValues(t, tt.expectedOptions, hc.options)
		})
	}
}

func TestDoWithNoLimit(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Send response to be tested
		rw.Write([]byte(`OK`))
		rw.WriteHeader(http.StatusOK)
	}))
	// Close the server when test finishes
	defer server.Close()

	httpClient := NewClient()
	httpClient.client = server.Client()

	for i := 0; i < 5; i++ {
		request, err := http.NewRequest(http.MethodGet, server.URL, nil)
		assert.NoError(t, err)

		response, err := httpClient.Do(request)
		assert.NoError(t, err)

		body, err := io.ReadAll(response.Body)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Equal(t, []byte(`OK`), body)
	}
}

func TestDoWithLimit_NonBlocking(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Send response to be tested
		rw.Write([]byte(`OK`))
		rw.WriteHeader(http.StatusOK)
	}))
	// Close the server when test finishes
	defer server.Close()

	httpClient := NewClient(WithLimitInterval(time.Second*2), WithLimit(5))
	httpClient.client = server.Client()

	for i := 0; i < 6; i++ {
		request, err := http.NewRequest(http.MethodGet, server.URL, nil)
		assert.NoError(t, err)

		if i < 5 {
			response, err := httpClient.Do(request)
			assert.NoError(t, err)

			body, err := io.ReadAll(response.Body)
			assert.NoError(t, err)

			assert.Equal(t, http.StatusOK, response.StatusCode)
			assert.Equal(t, []byte(`OK`), body)
		} else {
			response, err := httpClient.Do(request)
			assert.Nil(t, response)
			assert.True(t, errors.Is(err, ErrorRateLimitExceeded))
		}
	}
}

func TestDoWithLimit_Blocking(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Send response to be tested
		rw.Write([]byte(`OK`))
		rw.WriteHeader(http.StatusOK)
	}))
	// Close the server when test finishes
	defer server.Close()

	httpClient := NewClient(WithBlocking(), WithLimitInterval(time.Second*1), WithLimit(5))
	httpClient.client = server.Client()

	for i := 0; i < 6; i++ {
		request, err := http.NewRequest(http.MethodGet, server.URL, nil)
		assert.NoError(t, err)

		response, err := httpClient.Do(request)
		assert.NoError(t, err)

		body, err := io.ReadAll(response.Body)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Equal(t, []byte(`OK`), body)
	}
}
