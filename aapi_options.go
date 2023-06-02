package emailvalidator

import (
	"net/url"
)

type AbstractAPIOption struct {
	baseURL    *url.URL
	apiVersion string
	rate       AARate
	blocking   bool
}

type AbstractAPIOptionFunc func(*AbstractAPIOption) error

// WithAbstractAPIRate is used to specific rate limit on http client
func WithAbstractAPIRate(rate AARate) AbstractAPIOptionFunc {
	return func(opts *AbstractAPIOption) error {
		opts.rate = rate
		return nil
	}
}

// WithAbstractAPIBlocking is used to direct rateLimiter to wait until rate limit interval ends, if rate limit is reached.
func WithAbstractAPIBlocking() AbstractAPIOptionFunc {
	return func(opts *AbstractAPIOption) error {
		opts.blocking = true
		return nil
	}
}

// WithAbstractAPIBaseURL is used to set base url of abstract api service
func WithAbstractAPIBaseURL(url *url.URL) AbstractAPIOptionFunc {
	return func(opts *AbstractAPIOption) error {
		if url == nil || url.Path == "" {
			return ErrEmptyBaseURL
		}
		opts.baseURL = url
		return nil
	}
}

// WithAbstractAPIVersion is used to set API version of abstract api
func WithAbstractAPIVersion(version string) AbstractAPIOptionFunc {
	return func(opts *AbstractAPIOption) error {
		if version == "" {
			return ErrEmptyBaseURL
		}
		opts.apiVersion = version
		return nil
	}
}

func ParseAbstractAPIOptions(options ...AbstractAPIOptionFunc) (*AbstractAPIOption, error) {
	opts := &AbstractAPIOption{
		rate:     AbstractAPIFree,
		blocking: false,
	}

	for _, o := range options {
		if err := o(opts); err != nil {
			return nil, err
		}
	}

	return opts, nil
}
