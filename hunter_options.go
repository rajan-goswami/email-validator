package emailvalidator

import (
	"net/url"
)

type HunterAPIOption struct {
	baseURL    *url.URL
	apiVersion string
	apiSubPath string
	rate       HunterRate
	blocking   bool
}

type HunterAPIOptionFunc func(*HunterAPIOption) error

// WithHunterAPIRate is used to specific rate limit on http client
func WithHunterAPIRate(rate HunterRate) HunterAPIOptionFunc {
	return func(opts *HunterAPIOption) error {
		opts.rate = rate
		return nil
	}
}

// WithHunterAPIBlocking is used to direct rateLimiter to wait until rate limit interval ends, if rate limit is reached.
func WithHunterAPIBlocking() HunterAPIOptionFunc {
	return func(opts *HunterAPIOption) error {
		opts.blocking = true
		return nil
	}
}

// WithHunterAPIBaseURL is used to set base url of hunter api service
func WithHunterAPIBaseURL(url *url.URL) HunterAPIOptionFunc {
	return func(opts *HunterAPIOption) error {
		if url == nil || url.Path == "" {
			return ErrEmptyBaseURL
		}
		opts.baseURL = url
		return nil
	}
}

// WithHunterAPIVersion is used to set API version of abstract api
func WithHunterAPIVersion(version string) HunterAPIOptionFunc {
	return func(opts *HunterAPIOption) error {
		if version == "" {
			return ErrEmptyBaseURL
		}
		opts.apiVersion = version
		return nil
	}
}

func ParseHunterAPIOptions(options ...HunterAPIOptionFunc) (*HunterAPIOption, error) {
	opts := &HunterAPIOption{
		rate:     HunterAPIRate,
		blocking: false,
	}

	for _, o := range options {
		o(opts)
	}

	return opts, nil
}
