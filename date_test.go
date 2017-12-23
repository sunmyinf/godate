package godate

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	d := New(2017, time.November, 3)
	if d.IsZero() || d.Year() != 2017 || d.Month() != time.November || d.Day() != 3 {
		t.Errorf("unexpected date from New = %v", d)
	}

	d = New(1970, time.January, 1)
	if d.IsZero() || d.Year() != 1970 || d.Month() != time.January || d.Day() != 1 {
		t.Errorf("unexpected date from New = %v", d)
	}

	d = New(1969, time.December, 31)
	if d.IsZero() || d.Year() != 1969 || d.Month() != time.December || d.Day() != 31 {
		t.Errorf("unexpected date from New = %v", d)
	}

	d = Date{}
	if d.Year() != 1 || d.Month() != time.January || d.Day() != 1 {
		t.Errorf("unexpected date = %v, expect zero date", d)
	}
}

func TestNewFromTime(t *testing.T) {
	now := time.Now()
	zeroTime := time.Time{}

	times := []time.Time{
		now,
		time.Date(1969, 12, 1, 0, 0, 1, 0, time.UTC),
		zeroTime,
	}

	for i, tm := range times {
		d := NewFromTime(tm)

		if d.Year() != tm.Year() || d.Month() != tm.Month() || d.Day() != tm.Day() {
			t.Errorf("idx=%d unexpected date from NewFromTime=%v, expected=%v", i, d, tm.Format(RFC3339))
		}
	}
}

func TestNewFromElapsedDays(t *testing.T) {
	d := NewFromElapsedDays(ElapsedDays(719162))
	if d.Year() != 1970 || d.Month() != time.January || d.Day() != 1 {
		t.Errorf("unexpected date from NewFromElapsedDays = %v", d)
	}
}

func TestToday(t *testing.T) {
	today := Today()
	if today.IsZero() {
		t.Errorf("unexpected Today is zero")
	}
}

func TestDate_ToTime(t *testing.T) {
	d := New(2017, time.November, 3)
	tm := d.ToTime()

	if tm.IsZero() || tm.Format(time.RFC3339) != "2017-11-03T00:00:00Z" {
		t.Errorf("unexpected time from Date.ToTime = %s", tm.Format(time.RFC3339))
	}

	d = New(1969, time.December, 31)
	tm = d.ToTime()

	if tm.IsZero() || tm.Format(time.RFC3339) != "1969-12-31T00:00:00Z" {
		t.Errorf("unexpected time from Date.ToTime = %s", tm.Format(time.RFC3339))
	}

	d = Date{}
	tm = d.ToTime()

	if !tm.IsZero() {
		t.Errorf("expected zero value of time, but got not zero value. %v", tm)
	}
}

func TestDate_YearDay(t *testing.T) {
	d := New(2017, 10, 27)
	if d.YearDay() != 300 {
		t.Errorf("expected 302 as the 2017-10-27's YearDay, but got %d", d.YearDay())
	}

	d = New(2016, 12, 31)
	if d.YearDay() != 366 {
		t.Errorf("expected %d as the 2016-12-31's YearDay, but got %d", 366, d.YearDay())
	}
}

func TestDate_IsZero(t *testing.T) {
	d := Date{}
	if !d.IsZero() {
		t.Errorf("expected IsZero() is true, but got %v", d.IsZero())
	}

	d = Today()
	if d.IsZero() {
		t.Errorf("expected IsZero() is false, but got %v", d.IsZero())
	}
}
