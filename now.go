package now

import (
	"errors"
	"regexp"
	"time"
)

// BeginningOfMinute beginning of minute
func (now *Now) BeginningOfMinute() time.Time {
	return now.Truncate(time.Minute)
}

// BeginningOfHour beginning of hour
func (now *Now) BeginningOfHour() time.Time {
	y, m, d := now.Date()
	return time.Date(y, m, d, now.Time.Hour(), 0, 0, 0, now.Time.Location())
}

// BeginningOfDay beginning of day
func (now *Now) BeginningOfDay() time.Time {
	y, m, d := now.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, now.Time.Location())
}

// BeginningOfWeek beginning of week
func (now *Now) BeginningOfWeek() time.Time {
	t := now.BeginningOfDay()
	weekday := int(t.Weekday())

	if now.WeekStartDay != time.Sunday {
		weekStartDayInt := int(now.WeekStartDay)

		if weekday < weekStartDayInt {
			weekday = weekday + 7 - weekStartDayInt
		} else {
			weekday = weekday - weekStartDayInt
		}
	}
	return t.AddDate(0, 0, -weekday)
}

// BeginningOfMonth beginning of month
func (now *Now) BeginningOfMonth() time.Time {
	y, m, _ := now.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, now.Location())
}

// BeginningOfQuarter beginning of quarter
func (now *Now) BeginningOfQuarter() time.Time {
	month := now.BeginningOfMonth()
	offset := (int(month.Month()) - 1) % 3
	return month.AddDate(0, -offset, 0)
}

// BeginningOfHalf beginning of half year
func (now *Now) BeginningOfHalf() time.Time {
	month := now.BeginningOfMonth()
	offset := (int(month.Month()) - 1) % 6
	return month.AddDate(0, -offset, 0)
}

// BeginningOfYear BeginningOfYear beginning of year
func (now *Now) BeginningOfYear() time.Time {
	y, _, _ := now.Date()
	return time.Date(y, time.January, 1, 0, 0, 0, 0, now.Location())
}

// EndOfMinute end of minute
func (now *Now) EndOfMinute() time.Time {
	return now.BeginningOfMinute().Add(time.Minute - time.Nanosecond)
}

// EndOfHour end of hour
func (now *Now) EndOfHour() time.Time {
	return now.BeginningOfHour().Add(time.Hour - time.Nanosecond)
}

// EndOfDay end of day
func (now *Now) EndOfDay() time.Time {
	y, m, d := now.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), now.Location())
}

// EndOfWeek end of week
func (now *Now) EndOfWeek() time.Time {
	return now.BeginningOfWeek().AddDate(0, 0, 7).Add(-time.Nanosecond)
}

// EndOfMonth end of month
func (now *Now) EndOfMonth() time.Time {
	return now.BeginningOfMonth().AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// EndOfQuarter end of quarter
func (now *Now) EndOfQuarter() time.Time {
	return now.BeginningOfQuarter().AddDate(0, 3, 0).Add(-time.Nanosecond)
}

// EndOfHalf end of half year
func (now *Now) EndOfHalf() time.Time {
	return now.BeginningOfHalf().AddDate(0, 6, 0).Add(-time.Nanosecond)
}

// EndOfYear end of year
func (now *Now) EndOfYear() time.Time {
	return now.BeginningOfYear().AddDate(1, 0, 0).Add(-time.Nanosecond)
}

// Monday monday
/*
func (now *Now) Monday() time.Time {
	t := now.BeginningOfDay()
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return t.AddDate(0, 0, -weekday+1)
}
*/

// Monday returns the Monday Date of the week specified by now
func (now *Now) Monday(strs ...string) time.Time {
	return now.weekdayHelper(time.Monday, strs...)
}

// weekdayHelper consolidates the common weekday logic
func (now *Now) weekdayHelper(targetWeekday time.Weekday, strs ...string) time.Time {
	var parseTime time.Time
	var err error
	if len(strs) > 0 {
		parseTime, err = now.Parse(strs...)
		if err != nil {
			panic(err)
		}
	} else {
		parseTime = now.BeginningOfDay()
	}

	weekday := int(parseTime.Weekday())
	if weekday == 0 {
		weekday = 7
	}

	targetDay := int(targetWeekday)
	if targetDay == 0 {
		targetDay = 7
	}

	return parseTime.AddDate(0, 0, -weekday+targetDay)
}

// Tuesday returns the Tuesday Date of the week specified by now
func (now *Now) Tuesday(strs ...string) time.Time {
	return now.weekdayHelper(time.Tuesday, strs...)
}

// Wednesday returns the Wednesday Date of the week specified by now
func (now *Now) Wednesday(strs ...string) time.Time {
	return now.weekdayHelper(time.Wednesday, strs...)
}

// Thursday returns the Thursday Date of the week specified by now
func (now *Now) Thursday(strs ...string) time.Time {
	return now.weekdayHelper(time.Thursday, strs...)
}

// Friday returns the Friday Date of the week specified by now
func (now *Now) Friday(strs ...string) time.Time {
	return now.weekdayHelper(time.Friday, strs...)
}

// Saturday returns the Saturday Date of the week specified by now
func (now *Now) Saturday(strs ...string) time.Time {
	return now.weekdayHelper(time.Saturday, strs...)
}

// Sunday returns the Sunday Date of the week specified by now
func (now *Now) Sunday(strs ...string) time.Time {
	return now.weekdayHelper(time.Sunday, strs...)
}

// EndOfSunday end of sunday
func (now *Now) EndOfSunday() time.Time {
	return New(now.Sunday()).EndOfDay()
}

// Quarter returns the yearly quarter
func (now *Now) Quarter() uint {
	return (uint(now.Month())-1)/3 + 1
}

// Add returns a new Now instance with added duration
func (now *Now) Add(d time.Duration) *Now {
	return New(now.Time.Add(d))
}

// AddDays returns a new Now instance with added days
func (now *Now) AddDays(days int) *Now {
	return New(now.Time.AddDate(0, 0, days))
}

// AddWeeks returns a new Now instance with added weeks
func (now *Now) AddWeeks(weeks int) *Now {
	return New(now.Time.AddDate(0, 0, weeks*7))
}

// AddMonths returns a new Now instance with added months
func (now *Now) AddMonths(months int) *Now {
	return New(now.Time.AddDate(0, months, 0))
}

// AddYears returns a new Now instance with added years
func (now *Now) AddYears(years int) *Now {
	return New(now.Time.AddDate(years, 0, 0))
}

// FormatDate returns formatted date string (YYYY-MM-DD)
func (now *Now) FormatDate() string {
	return now.Format(DateFormat)
}

// FormatTime returns formatted time string (HH:MM:SS)
func (now *Now) FormatTime() string {
	return now.Format(TimeFormat)
}

// FormatDateTime returns formatted datetime string (YYYY-MM-DD HH:MM:SS)
func (now *Now) FormatDateTime() string {
	return now.Format(DateTimeFormat)
}

// FormatISO8601 returns formatted ISO8601 string
func (now *Now) FormatISO8601() string {
	return now.Format(ISO8601Format)
}

// FormatUS returns US formatted date string (MM/DD/YYYY)
func (now *Now) FormatUS() string {
	return now.Format(USFormat)
}

// FormatEuropean returns European formatted date string (DD/MM/YYYY)
func (now *Now) FormatEuropean() string {
	return now.Format(EuropeanFormat)
}

func (now *Now) parseWithFormat(str string, location *time.Location) (t time.Time, err error) {
	for _, format := range now.TimeFormats {
		t, err = time.ParseInLocation(format, str, location)

		if err == nil {
			return
		}
	}
	err = errors.New("Can't parse string as time: " + str)
	return
}

var hasTimeRegexp = regexp.MustCompile(`(\s+|^\s*|T)\d{1,2}((:\d{1,2})*|((:\d{1,2}){2}\.(\d{3}|\d{6}|\d{9})))(\s*$|[Z+-])`) // match 15:04:05, 15:04:05.000, 15:04:05.000000 15, 2017-01-01 15:04, 2021-07-20T00:59:10Z, 2021-07-20T00:59:10+08:00, 2021-07-20T00:00:10-07:00 etc
var onlyTimeRegexp = regexp.MustCompile(`^\s*\d{1,2}((:\d{1,2})*|((:\d{1,2}){2}\.(\d{3}|\d{6}|\d{9})))\s*$`)                // match 15:04:05, 15, 15:04:05.000, 15:04:05.000000, etc

// Parse parses string to time
func (now *Now) Parse(strs ...string) (t time.Time, err error) {
	var (
		setCurrentTime  bool
		parseTime       []int
		currentLocation = now.Location()
		onlyTimeInStr   = true
		currentTime     = formatTimeToList(now.Time)
	)

	for _, str := range strs {
		hasTimeInStr := hasTimeRegexp.MatchString(str) // match 15:04:05, 15
		onlyTimeInStr = hasTimeInStr && onlyTimeInStr && onlyTimeRegexp.MatchString(str)
		if t, err = now.parseWithFormat(str, currentLocation); err == nil {
			location := t.Location()
			parseTime = formatTimeToList(t)

			for i, v := range parseTime {
				// Don't reset hour, minute, second if current time str including time
				if hasTimeInStr && i <= 3 {
					continue
				}

				// If value is zero, replace it with current time
				if v == 0 {
					if setCurrentTime {
						parseTime[i] = currentTime[i]
					}
				} else {
					setCurrentTime = true
				}

				// if current time only includes time, should change day, month to current time
				if onlyTimeInStr {
					if i == 4 || i == 5 {
						parseTime[i] = currentTime[i]
						continue
					}
				}
			}

			t = time.Date(parseTime[6], time.Month(parseTime[5]), parseTime[4], parseTime[3], parseTime[2], parseTime[1], parseTime[0], location)
			currentTime = formatTimeToList(t)
		}
	}
	return
}

// MustParse must parse string to time or it will panic
func (now *Now) MustParse(strs ...string) (t time.Time) {
	t, err := now.Parse(strs...)
	if err != nil {
		panic(err)
	}
	return t
}

// Between check time between the begin, end time or not
func (now *Now) Between(begin, end string) bool {
	beginTime := now.MustParse(begin)
	endTime := now.MustParse(end)
	return now.After(beginTime) && now.Before(endTime)
}

// IsToday checks if the given time is within today
func IsToday(t time.Time) bool {
	n := New(time.Now())
	beginningOfDay := n.BeginningOfDay()
	endOfDay := n.EndOfDay()
	return (t.Equal(beginningOfDay) || t.After(beginningOfDay)) && (t.Equal(endOfDay) || t.Before(endOfDay))
}

// IsTimestampToday checks if the given Unix timestamp (in seconds, 10 digits) is within today
func IsTimestampToday(timestamp float64) bool {
	t := time.Unix(int64(timestamp), int64((timestamp-float64(int64(timestamp)))*1e9))
	return IsToday(t)
}

// IsMillisecondTimestampToday checks if the given Unix timestamp (in milliseconds, 13 digits) is within today
func IsMillisecondTimestampToday(timestamp float64) bool {
	seconds := int64(timestamp / 1000)
	nanoseconds := int64((timestamp - float64(seconds*1000)) * 1e6)
	t := time.Unix(seconds, nanoseconds)
	return IsToday(t)
}

// IsWeekend checks if the given time is Saturday or Sunday
func IsWeekend(t time.Time) bool {
	weekday := t.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}

// IsWeekday checks if the given time is Monday through Friday
func IsWeekday(t time.Time) bool {
	return !IsWeekend(t)
}

// isSameDay checks if two times are on the same day
func isSameDay(t1, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

// IsYesterday checks if the given time is yesterday
func IsYesterday(t time.Time) bool {
	yesterday := time.Now().AddDate(0, 0, -1)
	return isSameDay(t, yesterday)
}

// IsTomorrow checks if the given time is tomorrow
func IsTomorrow(t time.Time) bool {
	tomorrow := time.Now().AddDate(0, 0, 1)
	return isSameDay(t, tomorrow)
}

// IsSameWeek checks if two times are in the same week
func IsSameWeek(t1, t2 time.Time) bool {
	w1 := With(t1)
	w2 := With(t2)
	return w1.BeginningOfWeek().Equal(w2.BeginningOfWeek())
}

// IsSameMonth checks if two times are in the same month
func IsSameMonth(t1, t2 time.Time) bool {
	return t1.Year() == t2.Year() && t1.Month() == t2.Month()
}

// IsSameYear checks if two times are in the same year
func IsSameYear(t1, t2 time.Time) bool {
	return t1.Year() == t2.Year()
}

// DaysUntil returns the number of days until the target time
func DaysUntil(target time.Time) int {
	now := time.Now()
	nowDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	targetDay := time.Date(target.Year(), target.Month(), target.Day(), 0, 0, 0, 0, target.Location())
	diff := targetDay.Sub(nowDay)
	return int(diff.Hours() / 24)
}

// DaysSince returns the number of days since the target time
func DaysSince(target time.Time) int {
	now := time.Now()
	nowDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	targetDay := time.Date(target.Year(), target.Month(), target.Day(), 0, 0, 0, 0, target.Location())
	diff := nowDay.Sub(targetDay)
	return int(diff.Hours() / 24)
}

// WeeksUntil returns the number of weeks until the target time
func WeeksUntil(target time.Time) int {
	return DaysUntil(target) / 7
}

// WeeksSince returns the number of weeks since the target time
func WeeksSince(target time.Time) int {
	return DaysSince(target) / 7
}

// MonthsUntil returns the approximate number of months until the target time
func MonthsUntil(target time.Time) int {
	now := time.Now()
	years := target.Year() - now.Year()
	months := int(target.Month()) - int(now.Month())
	return years*12 + months
}

// MonthsSince returns the approximate number of months since the target time
func MonthsSince(target time.Time) int {
	now := time.Now()
	years := now.Year() - target.Year()
	months := int(now.Month()) - int(target.Month())
	return years*12 + months
}
