package emailvalidator

import (
	"net/url"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_parseHunterAPIOptions(t *testing.T) {

	defURL, err := url.Parse(defaultHunterAPIURL)
	assert.NoError(t, err)

	tests := []struct {
		name            string
		options         []HunterAPIOptionFunc
		expectedOptions *HunterAPIOption
		expectedError   error
	}{
		{
			name:    "Should parse with default options",
			options: []HunterAPIOptionFunc{},
			expectedOptions: &HunterAPIOption{
				baseURL:    defURL,
				apiVersion: defaultHunterAPIVersion,
				blocking:   false,
				rate:       HunterAPIRate,
			},
		},
		{
			name: "Should parse with custom options",
			options: []HunterAPIOptionFunc{
				WithHunterAPIVersion("/v10"),
				WithHunterAPIBlocking(),
				WithHunterAPIRate(HunterRate{Interval: time.Minute, Limit: 50}),
			},
			expectedOptions: &HunterAPIOption{
				baseURL:    defURL,
				apiVersion: "/v10",
				blocking:   true,
				rate:       HunterRate{Interval: time.Minute, Limit: 50},
			},
		},
		{
			name: "Should return error if baseURL is empty",
			options: []HunterAPIOptionFunc{
				WithHunterAPIBaseURL(nil),
			},
			expectedOptions: nil,
			expectedError:   ErrEmptyBaseURL,
		},
		{
			name: "Should return error if apiVersion is empty",
			options: []HunterAPIOptionFunc{
				WithHunterAPIVersion(""),
			},
			expectedOptions: nil,
			expectedError:   ErrEmptyAPIVersion,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseHunterAPIOptions(tt.options...)
			assert.Equal(t, tt.expectedError, err)
			if !reflect.DeepEqual(got, tt.expectedOptions) {
				t.Errorf("parseHunterAPIOptions() = %v, want %v", got, tt.expectedOptions)
			}
		})
	}
}
