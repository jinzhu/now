package now

import "time"

// Clock
type Clock interface {
	// Now returns the current time
	Now() time.Time
}

// TimeClock a clock implemention of time.Time
type TimeClock struct { }

func (*TimeClock) Now() time.Time {
	return time.Now()
}
