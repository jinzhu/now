package now

import (
	"testing"
	"time"
)

var (
	format          = "2006-01-02 15:04:05.999999999"
	locationCaracas *time.Location
	locationBerlin  *time.Location
	timeCaracas     time.Time
)

func init() {
	var err error
	if locationCaracas, err = time.LoadLocation("America/Caracas"); err != nil {
		panic(err)
	}

	if locationBerlin, err = time.LoadLocation("Europe/Berlin"); err != nil {
		panic(err)
	}

	timeCaracas = time.Date(2016, 1, 1, 12, 10, 0, 0, locationCaracas)
}

func assertT(t *testing.T) func(time.Time, string, string) {
	return func(actual time.Time, expected string, msg string) {
		actualStr := actual.Format(format)
		if actualStr != expected {
			t.Errorf("Failed %s: actual: %v, expected: %v", msg, actualStr, expected)
		}
	}
}

func TestBeginningOf(t *testing.T) {
	assert := assertT(t)

	n := time.Date(2013, 11, 18, 17, 51, 49, 123456789, time.UTC)

	assert(With(n).BeginningOfMinute(), "2013-11-18 17:51:00", "BeginningOfMinute")

	WeekStartDay = time.Monday
	assert(With(n).BeginningOfWeek(), "2013-11-18 00:00:00", "BeginningOfWeek, FirstDayMonday")

	WeekStartDay = time.Tuesday
	assert(With(n).BeginningOfWeek(), "2013-11-12 00:00:00", "BeginningOfWeek, FirstDayTuesday")

	WeekStartDay = time.Wednesday
	assert(With(n).BeginningOfWeek(), "2013-11-13 00:00:00", "BeginningOfWeek, FirstDayWednesday")

	WeekStartDay = time.Thursday
	assert(With(n).BeginningOfWeek(), "2013-11-14 00:00:00", "BeginningOfWeek, FirstDayThursday")

	WeekStartDay = time.Friday
	assert(With(n).BeginningOfWeek(), "2013-11-15 00:00:00", "BeginningOfWeek, FirstDayFriday")

	WeekStartDay = time.Saturday
	assert(With(n).BeginningOfWeek(), "2013-11-16 00:00:00", "BeginningOfWeek, FirstDaySaturday")

	WeekStartDay = time.Sunday
	assert(With(n).BeginningOfWeek(), "2013-11-17 00:00:00", "BeginningOfWeek, FirstDaySunday")

	assert(With(n).BeginningOfHour(), "2013-11-18 17:00:00", "BeginningOfHour")

	// Truncate with hour bug
	assert(With(timeCaracas).BeginningOfHour(), "2016-01-01 12:00:00", "BeginningOfHour Caracas")

	assert(With(n).BeginningOfDay(), "2013-11-18 00:00:00", "BeginningOfDay")

	location, err := time.LoadLocation("Japan")
	if err != nil {
		t.Fatalf("Error loading location: %v", err)
	}
	beginningOfDay := time.Date(2015, 05, 01, 0, 0, 0, 0, location)
	assert(With(beginningOfDay).BeginningOfDay(), "2015-05-01 00:00:00", "BeginningOfDay")

	// DST
	dstBeginningOfDay := time.Date(2017, 10, 29, 10, 0, 0, 0, locationBerlin)
	assert(With(dstBeginningOfDay).BeginningOfDay(), "2017-10-29 00:00:00", "BeginningOfDay DST")

	assert(With(n).BeginningOfWeek(), "2013-11-17 00:00:00", "BeginningOfWeek")

	dstBegginingOfWeek := time.Date(2017, 10, 30, 12, 0, 0, 0, locationBerlin)
	assert(With(dstBegginingOfWeek).BeginningOfWeek(), "2017-10-29 00:00:00", "BeginningOfWeek")

	dstBegginingOfWeek = time.Date(2017, 10, 29, 12, 0, 0, 0, locationBerlin)
	assert(With(dstBegginingOfWeek).BeginningOfWeek(), "2017-10-29 00:00:00", "BeginningOfWeek")

	WeekStartDay = time.Monday
	assert(With(n).BeginningOfWeek(), "2013-11-18 00:00:00", "BeginningOfWeek, FirstDayMonday")
	dstBegginingOfWeek = time.Date(2017, 10, 24, 12, 0, 0, 0, locationBerlin)
	assert(With(dstBegginingOfWeek).BeginningOfWeek(), "2017-10-23 00:00:00", "BeginningOfWeek, FirstDayMonday")

	dstBegginingOfWeek = time.Date(2017, 10, 29, 12, 0, 0, 0, locationBerlin)
	assert(With(dstBegginingOfWeek).BeginningOfWeek(), "2017-10-23 00:00:00", "BeginningOfWeek, FirstDayMonday")

	WeekStartDay = time.Sunday

	assert(With(n).BeginningOfMonth(), "2013-11-01 00:00:00", "BeginningOfMonth")

	// DST
	dstBeginningOfMonth := time.Date(2017, 10, 31, 0, 0, 0, 0, locationBerlin)
	assert(With(dstBeginningOfMonth).BeginningOfMonth(), "2017-10-01 00:00:00", "BeginningOfMonth DST")

	assert(With(n).BeginningOfQuarter(), "2013-10-01 00:00:00", "BeginningOfQuarter")

	// DST
	assert(With(dstBeginningOfMonth).BeginningOfQuarter(), "2017-10-01 00:00:00", "BeginningOfQuarter DST")
	dstBeginningOfQuarter := time.Date(2017, 11, 24, 0, 0, 0, 0, locationBerlin)
	assert(With(dstBeginningOfQuarter).BeginningOfQuarter(), "2017-10-01 00:00:00", "BeginningOfQuarter DST")

	assert(With(dstBeginningOfQuarter).BeginningOfHalf(), "2017-07-01 00:00:00", "BeginningOfHalf DST")

	assert(With(n.AddDate(0, -1, 0)).BeginningOfQuarter(), "2013-10-01 00:00:00", "BeginningOfQuarter")

	assert(With(n.AddDate(0, 1, 0)).BeginningOfQuarter(), "2013-10-01 00:00:00", "BeginningOfQuarter")

	assert(With(n.AddDate(0, 1, 0)).BeginningOfHalf(), "2013-07-01 00:00:00", "BeginningOfHalf")

	// DST
	assert(With(dstBeginningOfQuarter).BeginningOfYear(), "2017-01-01 00:00:00", "BeginningOfYear DST")

	assert(With(timeCaracas).BeginningOfYear(), "2016-01-01 00:00:00", "BeginningOfYear Caracas")
}

func TestEndOf(t *testing.T) {
	assert := assertT(t)

	n := time.Date(2013, 11, 18, 17, 51, 49, 123456789, time.UTC)

	assert(With(n).EndOfMinute(), "2013-11-18 17:51:59.999999999", "EndOfMinute")

	assert(With(n).EndOfHour(), "2013-11-18 17:59:59.999999999", "EndOfHour")

	assert(With(timeCaracas).EndOfHour(), "2016-01-01 12:59:59.999999999", "EndOfHour Caracas")

	assert(With(n).EndOfDay(), "2013-11-18 23:59:59.999999999", "EndOfDay")

	dstEndOfDay := time.Date(2017, 10, 29, 1, 0, 0, 0, locationBerlin)
	assert(With(dstEndOfDay).EndOfDay(), "2017-10-29 23:59:59.999999999", "EndOfDay DST")

	WeekStartDay = time.Tuesday
	assert(With(n).EndOfWeek(), "2013-11-18 23:59:59.999999999", "EndOfWeek, FirstDayTuesday")

	WeekStartDay = time.Wednesday
	assert(With(n).EndOfWeek(), "2013-11-19 23:59:59.999999999", "EndOfWeek, FirstDayWednesday")

	WeekStartDay = time.Thursday
	assert(With(n).EndOfWeek(), "2013-11-20 23:59:59.999999999", "EndOfWeek, FirstDayThursday")

	WeekStartDay = time.Friday
	assert(With(n).EndOfWeek(), "2013-11-21 23:59:59.999999999", "EndOfWeek, FirstDayFriday")

	WeekStartDay = time.Saturday
	assert(With(n).EndOfWeek(), "2013-11-22 23:59:59.999999999", "EndOfWeek, FirstDaySaturday")

	WeekStartDay = time.Sunday
	assert(With(n).EndOfWeek(), "2013-11-23 23:59:59.999999999", "EndOfWeek, FirstDaySunday")

	WeekStartDay = time.Monday
	assert(With(n).EndOfWeek(), "2013-11-24 23:59:59.999999999", "EndOfWeek, FirstDayMonday")

	dstEndOfWeek := time.Date(2017, 10, 24, 12, 0, 0, 0, locationBerlin)
	assert(With(dstEndOfWeek).EndOfWeek(), "2017-10-29 23:59:59.999999999", "EndOfWeek, FirstDayMonday")

	dstEndOfWeek = time.Date(2017, 10, 29, 12, 0, 0, 0, locationBerlin)
	assert(With(dstEndOfWeek).EndOfWeek(), "2017-10-29 23:59:59.999999999", "EndOfWeek, FirstDayMonday")

	WeekStartDay = time.Sunday
	assert(With(n).EndOfWeek(), "2013-11-23 23:59:59.999999999", "EndOfWeek")

	dstEndOfWeek = time.Date(2017, 10, 29, 0, 0, 0, 0, locationBerlin)
	assert(With(dstEndOfWeek).EndOfWeek(), "2017-11-04 23:59:59.999999999", "EndOfWeek")

	dstEndOfWeek = time.Date(2017, 10, 29, 12, 0, 0, 0, locationBerlin)
	assert(With(dstEndOfWeek).EndOfWeek(), "2017-11-04 23:59:59.999999999", "EndOfWeek")

	assert(With(n).EndOfMonth(), "2013-11-30 23:59:59.999999999", "EndOfMonth")

	assert(With(n).EndOfQuarter(), "2013-12-31 23:59:59.999999999", "EndOfQuarter")

	assert(With(n).EndOfHalf(), "2013-12-31 23:59:59.999999999", "EndOfHalf")

	assert(With(n.AddDate(0, -1, 0)).EndOfQuarter(), "2013-12-31 23:59:59.999999999", "EndOfQuarter")

	assert(With(n.AddDate(0, 1, 0)).EndOfQuarter(), "2013-12-31 23:59:59.999999999", "EndOfQuarter")

	assert(With(n.AddDate(0, 1, 0)).EndOfHalf(), "2013-12-31 23:59:59.999999999", "EndOfHalf")

	assert(With(n).EndOfYear(), "2013-12-31 23:59:59.999999999", "EndOfYear")

	n1 := time.Date(2013, 02, 18, 17, 51, 49, 123456789, time.UTC)
	assert(With(n1).EndOfMonth(), "2013-02-28 23:59:59.999999999", "EndOfMonth for 2013/02")

	n2 := time.Date(1900, 02, 18, 17, 51, 49, 123456789, time.UTC)
	assert(With(n2).EndOfMonth(), "1900-02-28 23:59:59.999999999", "EndOfMonth")
}

func TestMondayAndSunday(t *testing.T) {
	assert := assertT(t)

	n := time.Date(2013, 11, 19, 17, 51, 49, 123456789, time.UTC)
	n2 := time.Date(2013, 11, 24, 17, 51, 49, 123456789, time.UTC)
	nDst := time.Date(2017, 10, 29, 10, 0, 0, 0, locationBerlin)

	assert(With(n).Monday(), "2013-11-18 00:00:00", "Monday")

	assert(With(n2).Monday(), "2013-11-18 00:00:00", "Monday")

	assert(With(timeCaracas).Monday(), "2015-12-28 00:00:00", "Monday Caracas")

	assert(With(nDst).Monday(), "2017-10-23 00:00:00", "Monday DST")

	assert(With(n).Monday("17:51:49"), "2013-11-18 17:51:49", "Monday")

	assert(With(n).Monday("17:51"), "2013-11-18 17:51:00", "Monday")

	assert(With(n).Sunday(), "2013-11-24 00:00:00", "Sunday")

	assert(With(n).Sunday("18:19:20"), "2013-11-24 18:19:20", "Sunday")

	assert(With(n).Sunday("18:19"), "2013-11-24 18:19:00", "Sunday")

	assert(With(n2).Sunday(), "2013-11-24 00:00:00", "Sunday")

	assert(With(timeCaracas).Sunday(), "2016-01-03 00:00:00", "Sunday Caracas")

	assert(With(nDst).Sunday(), "2017-10-29 00:00:00", "Sunday DST")

	assert(With(n).EndOfSunday(), "2013-11-24 23:59:59.999999999", "EndOfSunday")

	assert(With(timeCaracas).EndOfSunday(), "2016-01-03 23:59:59.999999999", "EndOfSunday Caracas")

	assert(With(nDst).EndOfSunday(), "2017-10-29 23:59:59.999999999", "EndOfSunday DST")

	assert(With(n).BeginningOfWeek(), "2013-11-17 00:00:00", "BeginningOfWeek, FirstDayMonday")

	WeekStartDay = time.Monday
	assert(With(n).BeginningOfWeek(), "2013-11-18 00:00:00", "BeginningOfWeek, FirstDayMonday")
}

func TestParse(t *testing.T) {
	assert := assertT(t)

	n := time.Date(2013, 11, 18, 17, 51, 49, 123456789, time.UTC)

	assert(With(n).MustParse("2002"), "2002-01-01 00:00:00", "Parse 2002")

	assert(With(n).MustParse("2002-10"), "2002-10-01 00:00:00", "Parse 2002-10")

	assert(With(n).MustParse("2002-10-12"), "2002-10-12 00:00:00", "Parse 2002-10-12")

	assert(With(n).MustParse("2002-10-12 22"), "2002-10-12 22:00:00", "Parse 2002-10-12 22")

	assert(With(n).MustParse("2002-10-12 22:14"), "2002-10-12 22:14:00", "Parse 2002-10-12 22:14")

	assert(With(n).MustParse("2002-10-12 2:4"), "2002-10-12 02:04:00", "Parse 2002-10-12 2:4")

	assert(With(n).MustParse("2002-10-12 02:04"), "2002-10-12 02:04:00", "Parse 2002-10-12 02:04")

	assert(With(n).MustParse("2002-10-12 22:14:56"), "2002-10-12 22:14:56", "Parse 2002-10-12 22:14:56")

	assert(With(n).MustParse("2002-10-12 00:14:56"), "2002-10-12 00:14:56", "Parse 2002-10-12 00:14:56")

	assert(With(n).MustParse("2013-12-19 23:28:09.999999999 +0800 CST"), "2013-12-19 23:28:09.999999999", "Parse two strings 2013-12-19 23:28:09.999999999 +0800 CST")

	assert(With(n).MustParse("10-12"), "2013-10-12 00:00:00", "Parse 10-12")

	assert(With(n).MustParse("18"), "2013-11-18 18:00:00", "Parse 18 as hour")

	assert(With(n).MustParse("18:20"), "2013-11-18 18:20:00", "Parse 18:20")

	assert(With(n).MustParse("00:01"), "2013-11-18 00:01:00", "Parse 00:01")

	assert(With(n).MustParse("00:00:00"), "2013-11-18 00:00:00", "Parse 00:00:00")

	assert(With(n).MustParse("18:20:39"), "2013-11-18 18:20:39", "Parse 18:20:39")

	assert(With(n).MustParse("18:20:39", "2011-01-01"), "2011-01-01 18:20:39", "Parse two strings 18:20:39, 2011-01-01")

	assert(With(n).MustParse("2011-1-1", "18:20:39"), "2011-01-01 18:20:39", "Parse two strings 2011-01-01, 18:20:39")

	assert(With(n).MustParse("2011-01-01", "18"), "2011-01-01 18:00:00", "Parse two strings 2011-01-01, 18")

	assert(With(n).MustParse("2002-10-12T00:14:56Z"), "2002-10-12 00:14:56", "Parse 2002-10-12T00:14:56Z")
	assert(With(n).MustParse("2002-10-12T00:00:56Z"), "2002-10-12 00:00:56", "Parse 2002-10-12T00:00:56Z")
	assert(With(n).MustParse("2002-10-12T00:00:00.999Z"), "2002-10-12 00:00:00.999", "Parse 2002-10-12T00:00:00.999Z")
	assert(With(n).MustParse("2002-10-12T00:14:56.999999Z"), "2002-10-12 00:14:56.999999", "Parse 2002-10-12T00:14:56.999999Z")
	assert(With(n).MustParse("2002-10-12T00:00:56.999999999Z"), "2002-10-12 00:00:56.999999999", "Parse 2002-10-12T00:00:56.999999999Z")

	assert(With(n).MustParse("2002-10-12T00:14:56+08:00"), "2002-10-12 00:14:56", "Parse 2002-10-12T00:14:56+08:00")
	_, off := With(n).MustParse("2002-10-12T00:14:56+08:00").Zone()
	if off != 28800 {
		t.Errorf("Parse 2002-10-12T00:14:56+08:00 shouldn't lose time zone offset")
	}
	assert(With(n).MustParse("2002-10-12T00:00:56-07:00"), "2002-10-12 00:00:56", "Parse 2002-10-12T00:00:56-07:00")
	_, off2 := With(n).MustParse("2002-10-12T00:00:56-07:00").Zone()
	if off2 != -25200 {
		t.Errorf("Parse 2002-10-12T00:00:56-07:00 shouldn't lose time zone offset")
	}
	assert(With(n).MustParse("2002-10-12T00:01:12.333+0200"), "2002-10-12 00:01:12.333", "Parse 2002-10-12T00:01:12.333+0200")
	_, off3 := With(n).MustParse("2002-10-12T00:01:12.333+0200").Zone()
	if off3 != 7200 {
		t.Errorf("Parse 2002-10-12T00:01:12.333+0200 shouldn't lose time zone offset")
	}
	assert(With(n).MustParse("2002-10-12T00:00:56.999999999+08:00"), "2002-10-12 00:00:56.999999999", "Parse 2002-10-12T00:00:56.999999999+08:00")
	_, off4 := With(n).MustParse("2002-10-12T00:14:56.999999999+08:00").Zone()
	if off4 != 28800 {
		t.Errorf("Parse 2002-10-12T00:14:56.999999999+08:00 shouldn't lose time zone offset")
	}
	assert(With(n).MustParse("2002-10-12T00:00:56.666666-07:00"), "2002-10-12 00:00:56.666666", "Parse 2002-10-12T00:00:56.666666-07:00")
	_, off5 := With(n).MustParse("2002-10-12T00:00:56.666666-07:00").Zone()
	if off5 != -25200 {
		t.Errorf("Parse 2002-10-12T00:00:56.666666-07:00 shouldn't lose time zone offset")
	}
	assert(With(n).MustParse("2002-10-12T00:01:12.999999999-06"), "2002-10-12 00:01:12.999999999", "Parse 2002-10-12T00:01:12.999999999-06")
	_, off6 := With(n).MustParse("2002-10-12T00:01:12.999999999-06").Zone()
	if off6 != -21600 {
		t.Errorf("Parse 2002-10-12T00:01:12.999999999-06 shouldn't lose time zone offset")
	}

	TimeFormats = append(TimeFormats, "02 Jan 15:04")
	assert(With(n).MustParse("04 Feb 12:09"), "2013-02-04 12:09:00", "Parse 04 Feb 12:09 with specified format")

	assert(With(n).MustParse("23:28:9 Dec 19, 2013 PST"), "2013-12-19 23:28:09", "Parse 23:28:9 Dec 19, 2013 PST")

	if With(n).MustParse("23:28:9 Dec 19, 2013 PST").Location().String() != "PST" {
		t.Errorf("Parse 23:28:9 Dec 19, 2013 PST shouldn't lose time zone")
	}

	n2 := With(n).MustParse("23:28:9 Dec 19, 2013 PST")
	if With(n2).MustParse("10:20").Location().String() != "PST" {
		t.Errorf("Parse 10:20 shouldn't change time zone")
	}

	TimeFormats = append(TimeFormats, "2006-01-02T15:04:05.0")
	if MustParseInLocation(time.UTC, "2018-02-13T15:17:06.0").String() != "2018-02-13 15:17:06 +0000 UTC" {
		t.Errorf("ParseInLocation 2018-02-13T15:17:06.0")
	}

	TimeFormats = append(TimeFormats, "2006-01-02 15:04:05.000")
	assert(With(n).MustParse("2018-04-20 21:22:23.473"), "2018-04-20 21:22:23.473", "Parse 2018/04/20 21:22:23.473")

	TimeFormats = append(TimeFormats, "15:04:05.000")
	assert(With(n).MustParse("13:00:01.365"), "2013-11-18 13:00:01.365", "Parse 13:00:01.365")
}

func TestIsToday(t *testing.T) {
	now := time.Now()
	if !IsToday(now) {
		t.Error("Current time should be today")
	}

	yesterday := now.AddDate(0, 0, -1)
	if IsToday(yesterday) {
		t.Error("Yesterday should not be today")
	}

	tomorrow := now.AddDate(0, 0, 1)
	if IsToday(tomorrow) {
		t.Error("Tomorrow should not be today")
	}

	// Test with specific time in today
	todayMorning := time.Date(now.Year(), now.Month(), now.Day(), 6, 0, 0, 0, now.Location())
	if !IsToday(todayMorning) {
		t.Error("This morning should be today")
	}

	todayEvening := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	if !IsToday(todayEvening) {
		t.Error("This evening should be today")
	}
}

func TestIsWeekend(t *testing.T) {
	// Saturday
	saturday := time.Date(2023, 1, 7, 12, 0, 0, 0, time.UTC)
	if !IsWeekend(saturday) {
		t.Error("Saturday should be weekend")
	}

	// Sunday
	sunday := time.Date(2023, 1, 8, 12, 0, 0, 0, time.UTC)
	if !IsWeekend(sunday) {
		t.Error("Sunday should be weekend")
	}

	// Monday
	monday := time.Date(2023, 1, 9, 12, 0, 0, 0, time.UTC)
	if IsWeekend(monday) {
		t.Error("Monday should not be weekend")
	}

	// Wednesday
	wednesday := time.Date(2023, 1, 11, 12, 0, 0, 0, time.UTC)
	if IsWeekend(wednesday) {
		t.Error("Wednesday should not be weekend")
	}

	// Friday
	friday := time.Date(2023, 1, 13, 12, 0, 0, 0, time.UTC)
	if IsWeekend(friday) {
		t.Error("Friday should not be weekend")
	}
}

func TestIsWeekday(t *testing.T) {
	// Monday should be weekday
	monday := time.Date(2023, 1, 9, 12, 0, 0, 0, time.UTC)
	if !IsWeekday(monday) {
		t.Error("Monday should be weekday")
	}

	// Saturday should not be weekday
	saturday := time.Date(2023, 1, 7, 12, 0, 0, 0, time.UTC)
	if IsWeekday(saturday) {
		t.Error("Saturday should not be weekday")
	}

	// Sunday should not be weekday
	sunday := time.Date(2023, 1, 8, 12, 0, 0, 0, time.UTC)
	if IsWeekday(sunday) {
		t.Error("Sunday should not be weekday")
	}
}

func TestIsYesterday(t *testing.T) {
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)

	if !IsYesterday(yesterday) {
		t.Error("Yesterday should be yesterday")
	}

	if IsYesterday(now) {
		t.Error("Today should not be yesterday")
	}

	tomorrow := now.AddDate(0, 0, 1)
	if IsYesterday(tomorrow) {
		t.Error("Tomorrow should not be yesterday")
	}
}

func TestIsTomorrow(t *testing.T) {
	now := time.Now()
	tomorrow := now.AddDate(0, 0, 1)

	if !IsTomorrow(tomorrow) {
		t.Error("Tomorrow should be tomorrow")
	}

	if IsTomorrow(now) {
		t.Error("Today should not be tomorrow")
	}

	yesterday := now.AddDate(0, 0, -1)
	if IsTomorrow(yesterday) {
		t.Error("Yesterday should not be tomorrow")
	}
}

func TestIsSameWeek(t *testing.T) {
	// Monday and Wednesday in same week
	monday := time.Date(2023, 1, 9, 12, 0, 0, 0, time.UTC)
	wednesday := time.Date(2023, 1, 11, 12, 0, 0, 0, time.UTC)

	if !IsSameWeek(monday, wednesday) {
		t.Error("Monday and Wednesday should be in same week")
	}

	// Monday and next Monday should not be in same week
	nextMonday := time.Date(2023, 1, 16, 12, 0, 0, 0, time.UTC)
	if IsSameWeek(monday, nextMonday) {
		t.Error("Monday and next Monday should not be in same week")
	}
}

func TestIsSameMonth(t *testing.T) {
	// Same month, different days
	date1 := time.Date(2023, 1, 5, 12, 0, 0, 0, time.UTC)
	date2 := time.Date(2023, 1, 25, 12, 0, 0, 0, time.UTC)

	if !IsSameMonth(date1, date2) {
		t.Error("Same month dates should be in same month")
	}

	// Different months
	date3 := time.Date(2023, 2, 5, 12, 0, 0, 0, time.UTC)
	if IsSameMonth(date1, date3) {
		t.Error("Different month dates should not be in same month")
	}

	// Different years, same month
	date4 := time.Date(2024, 1, 5, 12, 0, 0, 0, time.UTC)
	if IsSameMonth(date1, date4) {
		t.Error("Different year dates should not be in same month")
	}
}

func TestIsSameYear(t *testing.T) {
	// Same year
	date1 := time.Date(2023, 1, 5, 12, 0, 0, 0, time.UTC)
	date2 := time.Date(2023, 12, 25, 12, 0, 0, 0, time.UTC)

	if !IsSameYear(date1, date2) {
		t.Error("Same year dates should be in same year")
	}

	// Different years
	date3 := time.Date(2024, 1, 5, 12, 0, 0, 0, time.UTC)
	if IsSameYear(date1, date3) {
		t.Error("Different year dates should not be in same year")
	}
}

func TestDaysUntil(t *testing.T) {
	now := time.Now()
	tomorrow := now.AddDate(0, 0, 1)
	yesterday := now.AddDate(0, 0, -1)
	nextWeek := now.AddDate(0, 0, 7)

	if days := DaysUntil(tomorrow); days != 1 {
		t.Errorf("Days until tomorrow should be 1, got %d", days)
	}

	if days := DaysUntil(yesterday); days != -1 {
		t.Errorf("Days until yesterday should be -1, got %d", days)
	}

	if days := DaysUntil(nextWeek); days != 7 {
		t.Errorf("Days until next week should be 7, got %d", days)
	}
}

func TestDaysSince(t *testing.T) {
	now := time.Now()
	tomorrow := now.AddDate(0, 0, 1)
	yesterday := now.AddDate(0, 0, -1)
	lastWeek := now.AddDate(0, 0, -7)

	if days := DaysSince(yesterday); days != 1 {
		t.Errorf("Days since yesterday should be 1, got %d", days)
	}

	if days := DaysSince(tomorrow); days != -1 {
		t.Errorf("Days since tomorrow should be -1, got %d", days)
	}

	if days := DaysSince(lastWeek); days != 7 {
		t.Errorf("Days since last week should be 7, got %d", days)
	}
}

func TestWeekdayFunctions(t *testing.T) {
	assert := assertT(t)

	// Test Tuesday function
	n := time.Date(2023, 1, 10, 12, 0, 0, 0, time.UTC) // Tuesday
	assert(With(n).Tuesday(), "2023-01-10 00:00:00", "Tuesday from Tuesday")
	assert(With(n).Tuesday("15:30"), "2023-01-10 15:30:00", "Tuesday with time")

	// Test Wednesday function
	assert(With(n).Wednesday(), "2023-01-11 00:00:00", "Wednesday from Tuesday")
	assert(With(n).Wednesday("09:15"), "2023-01-11 09:15:00", "Wednesday with time")

	// Test Thursday function
	assert(With(n).Thursday(), "2023-01-12 00:00:00", "Thursday from Tuesday")
	assert(With(n).Thursday("14:45"), "2023-01-12 14:45:00", "Thursday with time")

	// Test Friday function
	assert(With(n).Friday(), "2023-01-13 00:00:00", "Friday from Tuesday")
	assert(With(n).Friday("17:00"), "2023-01-13 17:00:00", "Friday with time")

	// Test Saturday function
	assert(With(n).Saturday(), "2023-01-14 00:00:00", "Saturday from Tuesday")
	assert(With(n).Saturday("10:30"), "2023-01-14 10:30:00", "Saturday with time")
}

func TestFluentInterface(t *testing.T) {
	assert := assertT(t)

	n := time.Date(2023, 1, 15, 12, 0, 0, 0, time.UTC)
	now := With(n)

	// Test AddDays
	assert(now.AddDays(1).Time, "2023-01-16 12:00:00", "AddDays(1)")
	assert(now.AddDays(-1).Time, "2023-01-14 12:00:00", "AddDays(-1)")
	assert(now.AddDays(7).Time, "2023-01-22 12:00:00", "AddDays(7)")

	// Test AddWeeks
	assert(now.AddWeeks(1).Time, "2023-01-22 12:00:00", "AddWeeks(1)")
	assert(now.AddWeeks(-1).Time, "2023-01-08 12:00:00", "AddWeeks(-1)")

	// Test AddMonths
	assert(now.AddMonths(1).Time, "2023-02-15 12:00:00", "AddMonths(1)")
	assert(now.AddMonths(-1).Time, "2022-12-15 12:00:00", "AddMonths(-1)")

	// Test AddYears
	assert(now.AddYears(1).Time, "2024-01-15 12:00:00", "AddYears(1)")
	assert(now.AddYears(-1).Time, "2022-01-15 12:00:00", "AddYears(-1)")

	// Test chaining
	assert(now.AddDays(1).AddWeeks(1).Time, "2023-01-23 12:00:00", "AddDays(1).AddWeeks(1)")
}

func TestFormatMethods(t *testing.T) {
	n := time.Date(2023, 1, 15, 14, 30, 45, 0, time.UTC)
	now := With(n)

	if date := now.FormatDate(); date != "2023-01-15" {
		t.Errorf("FormatDate() = %s, want %s", date, "2023-01-15")
	}

	if timeStr := now.FormatTime(); timeStr != "14:30:45" {
		t.Errorf("FormatTime() = %s, want %s", timeStr, "14:30:45")
	}

	if dateTime := now.FormatDateTime(); dateTime != "2023-01-15 14:30:45" {
		t.Errorf("FormatDateTime() = %s, want %s", dateTime, "2023-01-15 14:30:45")
	}

	if us := now.FormatUS(); us != "01/15/2023" {
		t.Errorf("FormatUS() = %s, want %s", us, "01/15/2023")
	}

	if european := now.FormatEuropean(); european != "15/01/2023" {
		t.Errorf("FormatEuropean() = %s, want %s", european, "15/01/2023")
	}
}

func TestTimestampToday(t *testing.T) {
	now := time.Now()

	// Test Unix timestamp (seconds)
	timestamp := float64(now.Unix())
	if !IsTimestampToday(timestamp) {
		t.Error("Current Unix timestamp should be today")
	}

	// Test millisecond timestamp
	millisecondTimestamp := float64(now.UnixNano() / 1e6)
	if !IsMillisecondTimestampToday(millisecondTimestamp) {
		t.Error("Current millisecond timestamp should be today")
	}

	// Test yesterday timestamp
	yesterday := now.AddDate(0, 0, -1)
	yesterdayTimestamp := float64(yesterday.Unix())
	if IsTimestampToday(yesterdayTimestamp) {
		t.Error("Yesterday timestamp should not be today")
	}
}

func TestAdditionalParseFormats(t *testing.T) {
	assert := assertT(t)
	n := time.Date(2013, 11, 18, 17, 51, 49, 123456789, time.UTC)

	TimeFormats = append(TimeFormats, "2006-01-02 15:04:05.000000")
	assert(With(n).MustParse("2010-01-01 07:24:23.131384"), "2010-01-01 07:24:23.131384", "Parse 2010-01-01 07:24:23.131384")
	assert(With(n).MustParse("00:00:00.182736"), "2013-11-18 00:00:00.182736", "Parse 00:00:00.182736")

	n3 := MustParse("2017-12-11T10:25:49Z")
	if n3.Location() != time.UTC {
		t.Errorf("time location should be UTC, but got %v", n3.Location())
	}
}

func TestBetween(t *testing.T) {
	tm := time.Date(2015, 06, 30, 17, 51, 49, 123456789, time.Now().Location())
	if !With(tm).Between("23:28:9 Dec 19, 2013 PST", "23:28:9 Dec 19, 2015 PST") {
		t.Errorf("Between")
	}

	if !With(tm).Between("2015-05-12 12:20", "2015-06-30 17:51:50") {
		t.Errorf("Between")
	}
}

func TestConfig(t *testing.T) {
	assert := assertT(t)

	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		t.Errorf("load location for Asia/Shanghai should returns no error, but got %v", err)
	}

	myConfig := Config{
		WeekStartDay: time.Monday,
		TimeLocation: location,
		TimeFormats:  []string{"2006-01-02 15:04:05"},
	}

	n := time.Date(2013, 11, 18, 17, 51, 49, 123456789, time.Now().Location()) // // 2013-11-18 17:51:49.123456789 Mon
	assert(myConfig.With(n).BeginningOfWeek(), "2013-11-18 00:00:00", "BeginningOfWeek, FirstDayMonday")

	if result, _ := myConfig.Parse("2018-02-13 15:17:06"); result.String() != "2018-02-13 15:17:06 +0800 CST" {
		t.Errorf("ParseInLocation 2018-02-13T15:17:06.0, got %v", result)
	}

	if result := myConfig.MustParse("2018-02-13 15:17:06"); result.String() != "2018-02-13 15:17:06 +0800 CST" {
		t.Errorf("ParseInLocation 2018-02-13T15:17:06.0, got %v", result)
	}
}

func TestQuarter(t *testing.T) {
	type test struct {
		givenDate       time.Time
		expectedQuarter uint
	}

	tests := []test{
		{time.Date(2021, 6, 18, 0, 0, 0, 0, time.UTC), 2},
		{time.Date(2021, 7, 18, 0, 0, 0, 0, time.UTC), 3},
	}

	for _, tc := range tests {
		got := With(tc.givenDate).Quarter()
		if got != tc.expectedQuarter {
			t.Fatalf("Quarter %d expected, got %d", tc.expectedQuarter, got)
		}
	}
}

func Example() {
	time.Now() // 2013-11-18 17:51:49.123456789 Mon

	BeginningOfMinute() // 2013-11-18 17:51:00 Mon
	BeginningOfHour()   // 2013-11-18 17:00:00 Mon
	BeginningOfDay()    // 2013-11-18 00:00:00 Mon
	BeginningOfWeek()   // 2013-11-17 00:00:00 Sun

	WeekStartDay = time.Monday // Set Monday as first day
	BeginningOfWeek()          // 2013-11-18 00:00:00 Mon
	BeginningOfMonth()         // 2013-11-01 00:00:00 Fri
	BeginningOfQuarter()       // 2013-10-01 00:00:00 Tue
	BeginningOfYear()          // 2013-01-01 00:00:00 Tue

	EndOfMinute() // 2013-11-18 17:51:59.999999999 Mon
	EndOfHour()   // 2013-11-18 17:59:59.999999999 Mon
	EndOfDay()    // 2013-11-18 23:59:59.999999999 Mon
	EndOfWeek()   // 2013-11-23 23:59:59.999999999 Sat

	WeekStartDay = time.Monday // Set Monday as first day
	EndOfWeek()                // 2013-11-24 23:59:59.999999999 Sun
	EndOfMonth()               // 2013-11-30 23:59:59.999999999 Sat
	EndOfQuarter()             // 2013-12-31 23:59:59.999999999 Tue
	EndOfYear()                // 2013-12-31 23:59:59.999999999 Tue

	// Use another time
	t := time.Date(2013, 02, 18, 17, 51, 49, 123456789, time.UTC)
	With(t).EndOfMonth() // 2013-02-28 23:59:59.999999999 Thu

	Monday()        // 2013-11-18 00:00:00 Mon
	Monday("17:44") // 2013-11-18 17:44:00 Mon
	Sunday()        // 2013-11-24 00:00:00 Sun
	Sunday("17:44") // 2013-11-24 17:44:00 Sun
	EndOfSunday()   // 2013-11-24 23:59:59.999999999 Sun
}
