package internal

import "time"

// Option defines options for HTTP Client
type Option struct {
	limitInterval time.Duration
	limit         int

	// Controls rate limiting behaviors
	blocking bool
}

// WithLimitInterval is used to set minimum time interval between first request to a limit
func WithLimitInterval(interval time.Duration) func(*Option) {
	return func(opts *Option) {
		opts.limitInterval = interval
	}
}

// WithLimit is used to set request bursts limit per limit interval
func WithLimit(limit int) func(*Option) {
	return func(opts *Option) {
		opts.limit = limit
	}
}

// WithBlocking is used to direct rateLimiter to use Wait() blocking call, if rate limit is reached until limitInterval
func WithBlocking() func(*Option) {
	return func(opts *Option) {
		opts.blocking = true
	}
}

func parseOptions(options ...func(*Option)) *Option {
	opts := &Option{
		limitInterval: defaultInterval,
		limit:         defaultLimit,
		blocking:      false,
	}

	for _, o := range options {
		o(opts)
	}
	return opts
}
