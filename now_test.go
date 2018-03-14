package now

import (
	"testing"
	"time"
)

var format = "2006-01-02 15:04:05.999999999"

func TestBeginningOf(t *testing.T) {
	n := time.Date(2013, 11, 18, 17, 51, 49, 123456789, time.UTC)

	if New(n).BeginningOfMinute().Format(format) != "2013-11-18 17:51:00" {
		t.Errorf("BeginningOfMinute")
	}

	if New(n).BeginningOfHour().Format(format) != "2013-11-18 17:00:00" {
		t.Errorf("BeginningOfHour")
	}

	if New(n).BeginningOfDay().Format(format) != "2013-11-18 00:00:00" {
		t.Errorf("BeginningOfDay")
	}

	location, _ := time.LoadLocation("Japan")
	beginningOfDay := time.Date(2015, 05, 01, 0, 0, 0, 0, location)
	if New(beginningOfDay).BeginningOfDay().Format(format) != "2015-05-01 00:00:00" {
		t.Errorf("BeginningOfDay")
	}

	WeekStartDay = time.Monday
	if New(n).BeginningOfWeek().Format(format) != "2013-11-18 00:00:00" {
		t.Errorf("BeginningOfWeek, FirstDayMonday")
	}

	WeekStartDay = time.Tuesday
	if New(n).BeginningOfWeek().Format(format) != "2013-11-12 00:00:00" {
		t.Errorf("BeginningOfWeek, FirstDayTuesday")
	}

	WeekStartDay = time.Wednesday
	if New(n).BeginningOfWeek().Format(format) != "2013-11-13 00:00:00" {
		t.Errorf("BeginningOfWeek, FirstDayWednesday")
	}

	WeekStartDay = time.Thursday
	if New(n).BeginningOfWeek().Format(format) != "2013-11-14 00:00:00" {
		t.Errorf("BeginningOfWeek, FirstDayThursday")
	}

	WeekStartDay = time.Friday
	if New(n).BeginningOfWeek().Format(format) != "2013-11-15 00:00:00" {
		t.Errorf("BeginningOfWeek, FirstDayFriday")
	}

	WeekStartDay = time.Saturday
	if New(n).BeginningOfWeek().Format(format) != "2013-11-16 00:00:00" {
		t.Errorf("BeginningOfWeek, FirstDaySaturday")
	}

	WeekStartDay = time.Sunday
	if New(n).BeginningOfWeek().Format(format) != "2013-11-17 00:00:00" {
		t.Errorf("BeginningOfWeek, FirstDaySunday")
	}

	if New(n).BeginningOfMonth().Format(format) != "2013-11-01 00:00:00" {
		t.Errorf("BeginningOfMonth")
	}

	if New(n).BeginningOfQuarter().Format(format) != "2013-10-01 00:00:00" {
		t.Error("BeginningOfQuarter")
	}

	if New(n.AddDate(0, -1, 0)).BeginningOfQuarter().Format(format) != "2013-10-01 00:00:00" {
		t.Error("BeginningOfQuarter")
	}

	if New(n.AddDate(0, 1, 0)).BeginningOfQuarter().Format(format) != "2013-10-01 00:00:00" {
		t.Error("BeginningOfQuarter")
	}

	if New(n).BeginningOfYear().Format(format) != "2013-01-01 00:00:00" {
		t.Errorf("BeginningOfYear")
	}
}

func TestEndOf(t *testing.T) {
	n := time.Date(2013, 11, 18, 17, 51, 49, 123456789, time.UTC)

	if New(n).EndOfMinute().Format(format) != "2013-11-18 17:51:59.999999999" {
		t.Errorf("EndOfMinute")
	}

	if New(n).EndOfHour().Format(format) != "2013-11-18 17:59:59.999999999" {
		t.Errorf("EndOfHour")
	}

	if New(n).EndOfDay().Format(format) != "2013-11-18 23:59:59.999999999" {
		t.Errorf("EndOfDay")
	}

	WeekStartDay = time.Monday
	if New(n).EndOfWeek().Format(format) != "2013-11-24 23:59:59.999999999" {
		t.Errorf("EndOfWeek, FirstDayMonday")
	}

	WeekStartDay = time.Tuesday
	if New(n).EndOfWeek().Format(format) != "2013-11-18 23:59:59.999999999" {
		t.Errorf("EndOfWeek, FirstDayTuesday")
	}

	WeekStartDay = time.Wednesday
	if New(n).EndOfWeek().Format(format) != "2013-11-19 23:59:59.999999999" {
		t.Errorf("EndOfWeek, FirstDayWednesday")
	}

	WeekStartDay = time.Thursday
	if New(n).EndOfWeek().Format(format) != "2013-11-20 23:59:59.999999999" {
		t.Errorf("EndOfWeek, FirstDayThursday")
	}

	WeekStartDay = time.Friday
	if New(n).EndOfWeek().Format(format) != "2013-11-21 23:59:59.999999999" {
		t.Errorf("EndOfWeek, FirstDayFriday")
	}

	WeekStartDay = time.Saturday
	if New(n).EndOfWeek().Format(format) != "2013-11-22 23:59:59.999999999" {
		t.Errorf("EndOfWeek, FirstDaySaturday")
	}

	WeekStartDay = time.Sunday
	if New(n).EndOfWeek().Format(format) != "2013-11-23 23:59:59.999999999" {
		t.Errorf("EndOfWeek, FirstDaySunday")
	}

	if New(n).EndOfMonth().Format(format) != "2013-11-30 23:59:59.999999999" {
		t.Errorf("EndOfMonth")
	}

	if New(n).EndOfQuarter().Format(format) != "2013-12-31 23:59:59.999999999" {
		t.Errorf("EndOfQuarter")
	}

	if New(n.AddDate(0, -1, 0)).EndOfQuarter().Format(format) != "2013-12-31 23:59:59.999999999" {
		t.Errorf("EndOfQuarter")
	}

	if New(n.AddDate(0, 1, 0)).EndOfQuarter().Format(format) != "2013-12-31 23:59:59.999999999" {
		t.Errorf("EndOfQuarter")
	}

	if New(n).EndOfYear().Format(format) != "2013-12-31 23:59:59.999999999" {
		t.Errorf("EndOfYear")
	}

	n1 := time.Date(2013, 02, 18, 17, 51, 49, 123456789, time.UTC)
	if New(n1).EndOfMonth().Format(format) != "2013-02-28 23:59:59.999999999" {
		t.Errorf("EndOfMonth for 2013/02")
	}

	n2 := time.Date(1900, 02, 18, 17, 51, 49, 123456789, time.UTC)
	if New(n2).EndOfMonth().Format(format) != "1900-02-28 23:59:59.999999999" {
		t.Errorf("EndOfMonth")
	}
}

func TestMondayAndSunday(t *testing.T) {
	n := time.Date(2013, 11, 19, 17, 51, 49, 123456789, time.UTC)
	n2 := time.Date(2013, 11, 24, 17, 51, 49, 123456789, time.UTC)

	if New(n).Monday().Format(format) != "2013-11-18 00:00:00" {
		t.Errorf("Monday")
	}

	if New(n2).Monday().Format(format) != "2013-11-18 00:00:00" {
		t.Errorf("Monday")
	}

	if New(n).Sunday().Format(format) != "2013-11-24 00:00:00" {
		t.Errorf("Sunday")
	}

	if New(n2).Sunday().Format(format) != "2013-11-24 00:00:00" {
		t.Errorf("Sunday")
	}

	if New(n).EndOfSunday().Format(format) != "2013-11-24 23:59:59.999999999" {
		t.Errorf("Sunday")
	}

	if New(n).BeginningOfWeek().Format(format) != "2013-11-17 00:00:00" {
		t.Errorf("BeginningOfWeek, FirstDayMonday")
	}

	WeekStartDay = time.Monday
	if New(n).BeginningOfWeek().Format(format) != "2013-11-18 00:00:00" {
		t.Errorf("BeginningOfWeek, FirstDayMonday")
	}
}

func TestParse(t *testing.T) {
	n := time.Date(2013, 11, 18, 17, 51, 49, 123456789, time.UTC)
	if New(n).MustParse("10-12").Format(format) != "2013-10-12 00:00:00" {
		t.Errorf("Parse 10-12")
	}

	if New(n).MustParse("2013-12-19 23:28:09.999999999 +0800 CST").Format(format) != "2013-12-19 23:28:09" {
		t.Errorf("Parse two strings 2013-12-19 23:28:09.999999999 +0800 CST")
	}

	if New(n).MustParse("2002-10-12 22:14").Format(format) != "2002-10-12 22:14:00" {
		t.Errorf("Parse 2002-10-12 22:14")
	}

	if New(n).MustParse("2002-10-12 2:4").Format(format) != "2002-10-12 02:04:00" {
		t.Errorf("Parse 2002-10-12 2:4")
	}

	if New(n).MustParse("2002-10-12 02:04").Format(format) != "2002-10-12 02:04:00" {
		t.Errorf("Parse 2002-10-12 02:04")
	}

	if New(n).MustParse("2002-10-12 22:14:56").Format(format) != "2002-10-12 22:14:56" {
		t.Errorf("Parse 2002-10-12 22:14:56")
	}

	if New(n).MustParse("2002-10-12").Format(format) != "2002-10-12 00:00:00" {
		t.Errorf("Parse 2002-10-12")
	}

	if New(n).MustParse("18").Format(format) != "2013-11-18 18:00:00" {
		t.Errorf("Parse 18 as hour")
	}

	if New(n).MustParse("18:20").Format(format) != "2013-11-18 18:20:00" {
		t.Errorf("Parse 18:20")
	}

	if New(n).MustParse("00:01").Format(format) != "2013-11-18 00:01:00" {
		t.Errorf("Parse 00:01")
	}

	if New(n).MustParse("18:20:39").Format(format) != "2013-11-18 18:20:39" {
		t.Errorf("Parse 18:20:39")
	}

	if New(n).MustParse("18:20:39", "2011-01-01").Format(format) != "2011-01-01 18:20:39" {
		t.Errorf("Parse two strings 18:20:39, 2011-01-01")
	}

	if New(n).MustParse("2011-1-1", "18:20:39").Format(format) != "2011-01-01 18:20:39" {
		t.Errorf("Parse two strings 2011-01-01, 18:20:39")
	}

	if New(n).MustParse("2011-01-01", "18").Format(format) != "2011-01-01 18:00:00" {
		t.Errorf("Parse two strings 2011-01-01, 18")
	}

	TimeFormats = append(TimeFormats, "02 Jan 15:04")
	if New(n).MustParse("04 Feb 12:09").Format(format) != "2013-02-04 12:09:00" {
		t.Errorf("Parse 04 Feb 12:09 with specified format")
	}

	if New(n).MustParse("23:28:9 Dec 19, 2013 PST").Format(format) != "2013-12-19 23:28:09" {
		t.Errorf("Parse 23:28:9 Dec 19, 2013 PST")
	}

	if New(n).MustParse("23:28:9 Dec 19, 2013 PST").Location().String() != "PST" {
		t.Errorf("Parse 23:28:9 Dec 19, 2013 PST shouldn't lose time zone")
	}

	n2 := New(n).MustParse("23:28:9 Dec 19, 2013 PST")
	if New(n2).MustParse("10:20").Location().String() != "PST" {
		t.Errorf("Parse 10:20 shouldn't change time zone")
	}

	TimeFormats = append(TimeFormats, "2006-01-02T15:04:05.0")
	if MustParseInLocation(time.UTC, "2018-02-13T15:17:06.0").String() != "2018-02-13 15:17:06 +0000 UTC" {
		t.Errorf("ParseInLocation 2018-02-13T15:17:06.0")
	}
}

func TestBetween(t *testing.T) {
	tm := time.Date(2015, 06, 30, 17, 51, 49, 123456789, time.Now().Location())
	if !New(tm).Between("23:28:9 Dec 19, 2013 PST", "23:28:9 Dec 19, 2015 PST") {
		t.Errorf("Between")
	}

	if !New(tm).Between("2015-05-12 12:20", "2015-06-30 17:51:50") {
		t.Errorf("Between")
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
	New(t).EndOfMonth() // 2013-02-28 23:59:59.999999999 Thu

	Monday()      // 2013-11-18 00:00:00 Mon
	Sunday()      // 2013-11-24 00:00:00 Sun
	EndOfSunday() // 2013-11-24 23:59:59.999999999 Sun
}
