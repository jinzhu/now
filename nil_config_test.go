package now

import (
	"testing"
	"time"
)

func TestNilConfigWithAndParse(t *testing.T) {
	var cfg *Config
	n := cfg.With(time.Date(2013, 11, 18, 17, 51, 49, 0, time.UTC))
	if n == nil || n.Config == nil {
		t.Fatal("expected non-nil Now with default Config")
	}
	// BeginningOfDay must not panic
	got := n.BeginningOfDay()
	want := time.Date(2013, 11, 18, 0, 0, 0, 0, time.UTC)
	if !got.Equal(want) {
		t.Fatalf("BeginningOfDay: got %v want %v", got, want)
	}

	tm, err := cfg.Parse("2013-11-18")
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if tm.Year() != 2013 || tm.Month() != 11 || tm.Day() != 18 {
		t.Fatalf("Parse result: %v", tm)
	}
}
