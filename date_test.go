package godate

import (
	"testing"
	"time"
)

func TestToday(t *testing.T) {
	today := Today()
	if today.t.IsZero() {
		t.Errorf("unexported t field of date is zero")
	}
}

func TestDate_YearDay(t *testing.T) {
	d := New(2017, 10, 27, time.UTC)
	if d.YearDay() != 300 {
		t.Errorf("expected 302 as the 2017-10-27's YearDay, but got %d", d.YearDay())
	}

	d = New(2016, 12, 31, time.UTC)
	if d.YearDay() != int(daysPerLeapYear/Day) {
		t.Errorf("expected %d as the 2016-12-31's YearDay, but got %d", int(daysPerLeapYear/Day), d.YearDay())
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

func TestDate_Year(t *testing.T) {
	d := New(2017, 10, 27, time.UTC)
	if d.Year() != 2017 {
		t.Errorf("expected %d, but got %d", 2017, d.Year())
	}
}

func TestDate_Month(t *testing.T) {
	d := New(2017, 10, 27, time.UTC)
	if d.Month() != time.October {
		t.Errorf("expected %d, but got %d", time.October, d.Month())
	}
}

func TestDate_Day(t *testing.T) {
	d := New(2017, 10, 27, time.UTC)
	if d.Day() != 27 {
		t.Errorf("expected %d, but got %d", 27, d.Day())
	}
}
