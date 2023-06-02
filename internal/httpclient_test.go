package internal

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {

	tests := []struct {
		name            string
		options         []func(*Option)
		expectedOptions *Option
		expectClient    bool
	}{
		{
			name:    "Should create client with default options",
			options: []func(*Option){},
			expectedOptions: &Option{
				blocking:      false,
				limit:         defaultLimit,
				limitInterval: defaultInterval,
			},
			expectClient: true,
		},
		{
			name: "Should create client with custom options",
			options: []func(*Option){
				WithBlocking(),
				WithLimit(10),
				WithLimitInterval(time.Second),
			},
			expectedOptions: &Option{
				blocking:      true,
				limit:         10,
				limitInterval: time.Second,
			},
			expectClient: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hc := NewClient(tt.options...)
			if tt.expectClient {
				assert.NotNil(t, hc)
				assert.NotNil(t, hc.client)
				assert.NotNil(t, hc.rateLimiter)
			} else {
				assert.Nil(t, hc)
			}
			assert.EqualValues(t, tt.expectedOptions, hc.options)
		})
	}
}
