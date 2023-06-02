package emailvalidator

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseAbstractAPIOptions(t *testing.T) {

	defURL, err := url.Parse(defaultAbstractAPIURL)
	assert.NoError(t, err)

	tests := []struct {
		name            string
		options         []AbstractAPIOptionFunc
		expectedOptions *AbstractAPIOption
		expectedError   error
	}{
		{
			name:    "Should parse with default options",
			options: []AbstractAPIOptionFunc{},
			expectedOptions: &AbstractAPIOption{
				baseURL:    defURL,
				apiVersion: defaultAbstractAPIVersion,
				blocking:   false,
				rate:       AbstractAPIFree,
			},
		},
		{
			name: "Should parse with custom options",
			options: []AbstractAPIOptionFunc{
				WithAbstractAPIVersion("/v10"),
				WithAbstractAPIBlocking(),
				WithAbstractAPIRate(AbstractAPIBusiness),
			},
			expectedOptions: &AbstractAPIOption{
				baseURL:    defURL,
				apiVersion: "/v10",
				blocking:   true,
				rate:       AbstractAPIBusiness,
			},
		},
		{
			name: "Should return error if baseURL is empty",
			options: []AbstractAPIOptionFunc{
				WithAbstractAPIBaseURL(nil),
			},
			expectedOptions: nil,
			expectedError:   ErrEmptyBaseURL,
		},
		{
			name: "Should return error if apiVersion is empty",
			options: []AbstractAPIOptionFunc{
				WithAbstractAPIVersion(""),
			},
			expectedOptions: nil,
			expectedError:   ErrEmptyAPIVersion,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseAbstractAPIOptions(tt.options...)
			assert.Equal(t, tt.expectedError, err)
			if !reflect.DeepEqual(got, tt.expectedOptions) {
				t.Errorf("parseAbstractAPIOptions() = %v, want %v", got, tt.expectedOptions)
			}
		})
	}
}
