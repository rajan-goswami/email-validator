package internal

import (
	"reflect"
	"testing"
	"time"
)

func Test_parseOptions(t *testing.T) {
	tests := []struct {
		name            string
		options         []func(*Option)
		expectedOptions *Option
	}{
		{
			name: "Should be able to set options",
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
		},
		{
			name:    "Should apply default limits",
			options: []func(*Option){},
			expectedOptions: &Option{
				blocking:      false,
				limit:         defaultLimit,
				limitInterval: defaultInterval,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseOptions(tt.options...); !reflect.DeepEqual(got, tt.expectedOptions) {
				t.Errorf("parseOptions() = %v, want %v", got, tt.expectedOptions)
			}
		})
	}
}
