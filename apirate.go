package emailvalidator

import "time"

type APIRate struct {
	Interval time.Duration
	Limit    int
}

var (
	// AbstractAPIFree represents API Rate limit under free plan
	AbstractAPIFree = APIRate{Interval: time.Second, Limit: 1}

	// AbstractAPIStarter represents API Rate limit under starter plan
	AbstractAPIStarter = APIRate{Interval: time.Second, Limit: 3}

	// AbstractAPIStandard represents API Rate limit under standard plan
	AbstractAPIStandard = APIRate{Interval: time.Second, Limit: 10}

	// AbstractAPIBusiness represents API Rate limit under business plan
	AbstractAPIBusiness = APIRate{Interval: time.Second, Limit: 25}

	// AbstractAPIProfessional represents API Rate limit under professional plan
	AbstractAPIProfessional = APIRate{Interval: time.Second, Limit: 50}

	// AbstractAPIGrowth represents API Rate limit under growth plan
	AbstractAPIGrowth = APIRate{Interval: time.Second, Limit: 100}
)
