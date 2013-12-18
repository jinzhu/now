// Package now is a time toolkit for golang.
//
// More details README here: https://github.com/jinzhu/now
//
//  import "github.com/jinzhu/now"
//
//  now.BeginningOfMinute() // 2013-11-18 17:51:00 Mon
//  now.BeginningOfDay()    // 2013-11-18 00:00:00 Mon
//  now.EndOfDay()          // 2013-11-18 23:59:59.999999999 Mon
package now

import "time"

var FirstDayMonday bool

type Now struct {
	time.Time
}

func New(t time.Time) *Now {
	return &Now{t}
}

func BeginningOfMinute() time.Time {
	return New(time.Now()).BeginningOfMinute()
}

func BeginningOfHour() time.Time {
	return New(time.Now()).BeginningOfHour()
}

func BeginningOfDay() time.Time {
	return New(time.Now()).BeginningOfDay()
}

func BeginningOfWeek() time.Time {
	return New(time.Now()).BeginningOfWeek()
}

func BeginningOfMonth() time.Time {
	return New(time.Now()).BeginningOfMonth()
}

func BeginningOfYear() time.Time {
	return New(time.Now()).BeginningOfYear()
}

func EndOfMinute() time.Time {
	return New(time.Now()).EndOfMinute()
}

func EndOfHour() time.Time {
	return New(time.Now()).EndOfHour()
}

func EndOfDay() time.Time {
	return New(time.Now()).EndOfDay()
}

func EndOfWeek() time.Time {
	return New(time.Now()).EndOfWeek()
}

func EndOfMonth() time.Time {
	return New(time.Now()).EndOfMonth()
}

func EndOfYear() time.Time {
	return New(time.Now()).EndOfYear()
}

func NextOccurrenceOf(month time.Month, day int) time.Time {
	return New(time.Now()).NextOccurrenceOf(month, day)
}

func LastOccurrenceOf(month time.Month, day int) time.Time {
	return New(time.Now()).LastOccurrenceOf(month, day)
}

func Monday() time.Time {
	return New(time.Now()).EndOfYear()
}

func Sunday() time.Time {
	return New(time.Now()).EndOfYear()
}

func EndOfSunday() time.Time {
	return New(time.Now()).EndOfSunday()
}
