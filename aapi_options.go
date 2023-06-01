package emailvalidator

import (
	"net/url"
)

type AbstractAPIOption struct {
	baseURL    *url.URL
	apiVersion string
	rate       APIRate
	blocking   bool
}

type AbstractAPIOptionFunc func(*AbstractAPIOption) error

// WithAPIRate is used to specific rate limit on http client
func WithAPIRate(rate APIRate) AbstractAPIOptionFunc {
	return func(opts *AbstractAPIOption) error {
		opts.rate = rate
		return nil
	}
}

// WithBlocking is used to direct rateLimiter to wait until rate limit interval ends, if rate limit is reached.
func WithBlocking() AbstractAPIOptionFunc {
	return func(opts *AbstractAPIOption) error {
		opts.blocking = true
		return nil
	}
}

// WithBaseURL is used to set base url of abstract api service
func WithBaseURL(url *url.URL) AbstractAPIOptionFunc {
	return func(opts *AbstractAPIOption) error {
		if url == nil || url.Path == "" {
			return ErrEmptyBaseURL
		}
		opts.baseURL = url
		return nil
	}
}

// WithAPIVersion is used to set API version of abstract api
func WithAPIVersion(version string) AbstractAPIOptionFunc {
	return func(opts *AbstractAPIOption) error {
		if version == "" {
			return ErrEmptyBaseURL
		}
		opts.apiVersion = version
		return nil
	}
}

func ParseOptions(options ...AbstractAPIOptionFunc) (*AbstractAPIOption, error) {
	opts := &AbstractAPIOption{
		rate:     AbstractAPIFree,
		blocking: false,
	}

	for _, o := range options {
		o(opts)
	}

	return opts, nil
}
