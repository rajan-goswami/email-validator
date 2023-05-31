package internal

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

var ErrorRateLimitExceeded = fmt.Errorf("rate limit is exceeded")

const (
	DefaultInterval time.Duration = 0 * time.Second
	DefaultLimit    int           = -1
)

// HTTPClient defines HTTP client with rate limiting capability
type HTTPClient struct {
	client      *http.Client
	rateLimiter *rate.Limiter

	options *Option
}

// NewClient return http client with a rate limiter
func NewClient(options ...OptionFunc) *HTTPClient {
	opts := ParseOptions(options...)
	c := &HTTPClient{
		client:      http.DefaultClient,
		rateLimiter: rate.NewLimiter(rate.Every(opts.limitInterval), opts.limit),
		options:     opts,
	}
	return c
}

func (c *HTTPClient) Do(req *http.Request) (*http.Response, error) {
	ctx := context.Background()

	if c.options.blocking {
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
