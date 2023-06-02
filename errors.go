package emailvalidator

import "errors"

var (
	// ErrEmptyAPIKey is returned when api key is empty.
	ErrEmptyAPIKey = errors.New("empty api key provided")

	// ErrEmptyBaseURL is returned when baseURL is empty.
	ErrEmptyBaseURL = errors.New("empty baseURL")

	// ErrEmptyAPIVersion is returned when api version is empty.
	ErrEmptyAPIVersion = errors.New("empty api version")

	// ErrEmptyAPISubPath is returned when API sub path is empty in the case of hunter.io APIs.
	ErrEmptyAPISubPath = errors.New("empty api sub-path")
)
