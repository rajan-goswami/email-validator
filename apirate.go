package emailvalidator

import "time"

type apiRate struct {
	Interval time.Duration
	Limit    int
}

type AARate apiRate
type HunterRate apiRate

var (
	// AbstractAPIFree represents API Rate limit under free plan
	AbstractAPIFree = AARate{Interval: time.Second, Limit: 1}

	// AbstractAPIStarter represents API Rate limit under starter plan
	AbstractAPIStarter = AARate{Interval: time.Second, Limit: 3}

	// AbstractAPIStandard represents API Rate limit under standard plan
	AbstractAPIStandard = AARate{Interval: time.Second, Limit: 10}

	// AbstractAPIBusiness represents API Rate limit under business plan
	AbstractAPIBusiness = AARate{Interval: time.Second, Limit: 25}

	// AbstractAPIProfessional represents API Rate limit under professional plan
	AbstractAPIProfessional = AARate{Interval: time.Second, Limit: 50}

	// AbstractAPIGrowth represents API Rate limit under growth plan
	AbstractAPIGrowth = AARate{Interval: time.Second, Limit: 100}
)

var (
	// HunterAPIRate represents API Rate limit for email verify API
	HunterAPIRate = HunterRate{Interval: time.Second, Limit: 10}
)
