package emailvalidator

import "time"

type APIRate struct {
	Interval time.Duration
	Limit    int
}
