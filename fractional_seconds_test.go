package now

import (
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestParseFractionalSeconds(t *testing.T) {
	reference := time.Date(2013, 11, 18, 17, 51, 49, 987654321, time.UTC)
	for _, fraction := range []string{"0", "1", "12", "123", "1234", "12345", "123456", "1234567", "12345678", "123456789"} {
		nanoseconds, err := strconv.Atoi(fraction + strings.Repeat("0", 9-len(fraction)))
		if err != nil {
			t.Fatal(err)
		}
		for _, separator := range []string{".", ","} {
			suffix := separator + fraction
			for _, test := range []struct {
				input string
				want  time.Time
			}{
				{"2002-10-12T00:00:00" + suffix + "Z", time.Date(2002, 10, 12, 0, 0, 0, nanoseconds, time.UTC)},
				{"2002-10-12T00:00:00" + suffix + "+08:00", time.Date(2002, 10, 12, 0, 0, 0, nanoseconds, time.FixedZone("", 8*60*60))},
				{"2002-10-12 00:00:00" + suffix, time.Date(2002, 10, 12, 0, 0, 0, nanoseconds, time.UTC)},
				{"00:00:00" + suffix, time.Date(2013, 11, 18, 0, 0, 0, nanoseconds, time.UTC)},
			} {
				t.Run(test.input, func(t *testing.T) {
					got, err := With(reference).Parse(test.input)
					if err != nil {
						t.Fatal(err)
					}
					if !got.Equal(test.want) {
						t.Errorf("Parse(%q) = %s, want %s", test.input, got.Format(time.RFC3339Nano), test.want.Format(time.RFC3339Nano))
					}
					_, gotOffset := got.Zone()
					_, wantOffset := test.want.Zone()
					if gotOffset != wantOffset {
						t.Errorf("zone offset = %d, want %d", gotOffset, wantOffset)
					}
				})
			}
		}
	}
}
