package emailvalidator

import "errors"

var (
	ErrEmptyAPIKey     = errors.New("empty api key provided")
	ErrEmptyBaseURL    = errors.New("empty baseURL")
	ErrEmptyAPIVersion = errors.New("empty api version")
	ErrEmptyAPISubPath = errors.New("empty api sub-path")
)
