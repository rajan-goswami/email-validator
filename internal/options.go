package internal

import "time"

// Option defines options for HTTP Client
type Option struct {
	limitInterval time.Duration
	limit         int

	// Controls rate limiting behaviors
	blocking bool
}

type OptionFunc func(*Option)

// WithLimitInterval is used to set minimum time interval between first request to a limit
func WithLimitInterval(interval time.Duration) OptionFunc {
	return func(opts *Option) {
		opts.limitInterval = interval
	}
}

// WithLimit is used to set request bursts limit per limit interval
func WithLimit(limit int) OptionFunc {
	return func(opts *Option) {
		opts.limit = limit
	}
}

// WithBlocking is used to direct rateLimiter to use Wait() blocking call, if rate limit is reached until limitInterval
func WithBlocking() OptionFunc {
	return func(opts *Option) {
		opts.blocking = true
	}
}

func ParseOptions(options ...OptionFunc) *Option {
	opts := &Option{
		limitInterval: DefaultInterval,
		limit:         DefaultLimit,
		blocking:      false,
	}

	for _, o := range options {
		o(opts)
	}
	return opts
}
