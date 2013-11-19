package now

import "time"

func (now *Now) BeginningOfMinute() time.Time {
	return now.time.Truncate(time.Minute)
}

func (now *Now) BeginningOfHour() time.Time {
	return now.time.Truncate(time.Hour)
}

func (now *Now) BeginningOfDay() time.Time {
	d := time.Duration(-now.time.Hour()) * time.Hour
	return now.time.Truncate(time.Hour).Add(d)
}

func (now *Now) BeginningOfWeek() time.Time {
	t := now.BeginningOfDay()
	weekday := int(t.Weekday())
	if FirstDayMonday {
		if weekday == 0 {
			weekday = 7
		}
		weekday = weekday - 1
	}

	d := time.Duration(-weekday) * 24 * time.Hour
	return t.Truncate(time.Hour).Add(d)
}

func (now *Now) BeginningOfMonth() time.Time {
	t := now.BeginningOfDay()
	d := time.Duration(-int(t.Day())+1) * 24 * time.Hour
	return t.Truncate(time.Hour).Add(d)
}

func (now *Now) BeginningOfYear() time.Time {
	t := now.BeginningOfDay()
	d := time.Duration(-int(t.YearDay())+1) * 24 * time.Hour
	return t.Truncate(time.Hour).Add(d)
}

func (now *Now) EndOfMinute() time.Time {
	return now.BeginningOfMinute().Add(time.Minute - time.Nanosecond)
}

func (now *Now) EndOfHour() time.Time {
	return now.BeginningOfHour().Add(time.Hour - time.Nanosecond)
}

func (now *Now) EndOfDay() time.Time {
	return now.BeginningOfDay().Add(24*time.Hour - time.Nanosecond)
}

func (now *Now) EndOfWeek() time.Time {
	return now.BeginningOfWeek().AddDate(0, 0, 7).Add(-time.Nanosecond)
}

func (now *Now) EndOfMonth() time.Time {
	return now.BeginningOfMonth().AddDate(0, 1, 0).Add(-time.Nanosecond)
}

func (now *Now) EndOfYear() time.Time {
	return now.BeginningOfYear().AddDate(1, 0, 0).Add(-time.Nanosecond)
}

func (now *Now) Monday() time.Time {
	t := now.BeginningOfDay()
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	d := time.Duration(-weekday+1) * 24 * time.Hour
	return t.Truncate(time.Hour).Add(d)
}

func (now *Now) Sunday() time.Time {
	t := now.BeginningOfDay()
	weekday := int(t.Weekday())
	if weekday == 0 {
		return t
	} else {
		d := time.Duration(7-weekday) * 24 * time.Hour
		return t.Truncate(time.Hour).Add(d)
	}
}

func (now *Now) EndOfSunday() time.Time {
	return now.Sunday().Add(24*time.Hour - time.Nanosecond)
}
