package emailvalidator

import (
	"net/url"
)

// AbstractAPIOption specifies options for AbstractAPI client
type AbstractAPIOption struct {
	baseURL    *url.URL
	apiVersion string
	rate       AARate
	blocking   bool
}

// AbstractAPIOptionFunc is a function type for setting AbstractAPI client options
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
		if url == nil || url.String() == "" {
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
			return ErrEmptyAPIVersion
		}
		opts.apiVersion = version
		return nil
	}
}

func parseAbstractAPIOptions(options ...AbstractAPIOptionFunc) (*AbstractAPIOption, error) {
	defAPIURL, _ := url.Parse(defaultAbstractAPIURL)

	opts := &AbstractAPIOption{
		rate:       AbstractAPIFree,
		blocking:   false,
		baseURL:    defAPIURL,
		apiVersion: defaultAbstractAPIVersion,
	}

	for _, o := range options {
		if err := o(opts); err != nil {
			return nil, err
		}
	}

	return opts, nil
}
