package now

import "time"

func Example() {
	time.Now() // 2013-11-18 17:51:49.123456789 Mon

	BeginningOfMinute() // 2013-11-18 17:51:00 Mon
	BeginningOfHour()   // 2013-11-18 17:00:00 Mon
	BeginningOfDay()    // 2013-11-18 00:00:00 Mon
	BeginningOfWeek()   // 2013-11-17 00:00:00 Sun

	FirstDayMonday = true // Set Monday as first day
	BeginningOfWeek()     // 2013-11-18 00:00:00 Mon
	BeginningOfMonth()    // 2013-11-01 00:00:00 Fri
	BeginningOfYear()     // 2013-01-01 00:00:00 Tue

	EndOfMinute() // 2013-11-18 17:51:59.999999999 Mon
	EndOfHour()   // 2013-11-18 17:59:59.999999999 Mon
	EndOfDay()    // 2013-11-18 23:59:59.999999999 Mon
	EndOfWeek()   // 2013-11-23 23:59:59.999999999 Sat

	FirstDayMonday = true // Set Monday as first day
	EndOfWeek()           // 2013-11-24 23:59:59.999999999 Sun
	EndOfMonth()          // 2013-11-30 23:59:59.999999999 Sat
	EndOfYear()           // 2013-12-31 23:59:59.999999999 Tue

	// Use another time
	t := time.Date(2013, 02, 18, 17, 51, 49, 123456789, time.Now().Location())
	New(t).EndOfMonth() // 2013-02-28 23:59:59.999999999 Thu
}
