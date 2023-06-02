package internal

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

// ErrorRateLimitExceeded is returned when api rate limit is hit by HTTPClient
var ErrorRateLimitExceeded = fmt.Errorf("rate limit is exceeded")

const (
	defaultInterval time.Duration = 0 * time.Second
	defaultLimit    int           = -1
)

// HTTPClient defines HTTP client with rate limiting capability
type HTTPClient struct {
	client      *http.Client
	rateLimiter *rate.Limiter

	options *Option
}

// NewClient return http client with a rate limiter
func NewClient(options ...func(*Option)) *HTTPClient {
	opts := parseOptions(options...)
	c := &HTTPClient{
		client:      http.DefaultClient,
		rateLimiter: rate.NewLimiter(rate.Every(opts.limitInterval), opts.limit),
		options:     opts,
	}
	return c
}

// Do performs http request and returns response or error
// It applies client-side rate limits on outgoing API calls, based on the configured rate.
// If blocking is enabled and rate limit is reached, it will wait until rate limit interval ends while making api call.
// Otherwise it returns ErrorRateLimitExceeded
func (c *HTTPClient) Do(req *http.Request) (*http.Response, error) {
	if c.options.blocking {
		ctx := context.Background()
		err := c.rateLimiter.Wait(ctx) // This is a blocking call. Honors the rate limit
		if err != nil {
			return nil, err
		}
	} else {
		// Check if we are under the limit
		ok := c.rateLimiter.Allow()
		if !ok {
			return nil, errors.Unwrap(fmt.Errorf("HTTPClient: %w", ErrorRateLimitExceeded))
		}
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
