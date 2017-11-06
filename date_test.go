package godate

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	d := New(2017, time.November, 3)
	if d.IsZero() || d.Year != 2017 || d.Month != time.November || d.Day != 3 {
		t.Errorf("unexpected date from New = %v", d)
	}
}

func TestNewFromTime(t *testing.T) {
	tm := time.Now()
	d := NewFromTime(tm)

	if d.IsZero() || d.Year != tm.Year() || d.Month != tm.Month() || d.Day != tm.Day() {
		t.Errorf("unexpected date from NewFromTime = %v", d)
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
	time.Now()

	if tm.IsZero() || tm.Format(time.RFC3339) != "2017-11-03T00:00:00Z" {
		t.Errorf("unexpected time from Date.ToTime = %s", tm.Format(time.RFC3339))
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
