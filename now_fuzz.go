// +build gofuzz

package now

import (
	"time"
)

func FuzzParse(input []byte) int {
	_, err := Parse(string(input))
	if err != nil {
		return 0
	}
	return 1
}

// Maps bytes to API calls
func FuzzAPI(program []byte) int {
	t := New(time.Unix(1, 0))
	for _, step := range program {
		switch step {
		case 0:
			t = New(t.BeginningOfDay())
		case 1:
			t = New(t.BeginningOfHalf())
		case 2:
			t = New(t.BeginningOfHour())
		case 3:
			t = New(t.BeginningOfMinute())
		case 4:
			t = New(t.BeginningOfMonth())
		case 5:
			t = New(t.BeginningOfQuarter())
		case 6:
			t = New(t.BeginningOfWeek())
		case 7:
			t = New(t.BeginningOfYear())
		case 8:
			t = New(t.EndOfDay())
		case 9:
			t = New(t.EndOfHalf())
		case 10:
			t = New(t.EndOfHour())
		case 11:
			t = New(t.EndOfMinute())
		case 12:
			t = New(t.EndOfMonth())
		case 13:
			t = New(t.EndOfQuarter())
		case 14:
			t = New(t.EndOfSunday())
		case 15:
			t = New(t.EndOfWeek())
		case 16:
			t = New(t.EndOfYear())
		case 17:
			t = New(t.Monday())
		case 18:
			t = New(t.Sunday())
		default:
			return -1
		}
	}
	return 1
}
