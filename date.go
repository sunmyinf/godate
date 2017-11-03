package godate

import "time"

// Date has Year, Month, Day fields.
// Date don't have Location or TimeZone information.
type Date struct {
	Year  int
	Month time.Month
	Day   int
}

// New returns Date from specified year, month, day.
func New(year int, month time.Month, day int) Date {
	return Date{year, month, day}
}

// NewFromTime returns Date generated from time.Time's value
func NewFromTime(t time.Time) Date {
	return New(t.Year(), t.Month(), t.Day())
}

// Today returns Date of today
func Today() Date {
	now := time.Now()
	return Date{
		now.Year(),
		now.Month(),
		now.Day(),
	}
}

// ToTime returns time.Time instance in UTC from d.
func (d Date) ToTime() time.Time {
	return time.Date(d.Year, d.Month, d.Day, 0, 0, 0, 0, time.UTC)
}

// YearDay is wrapper of YearDay method of time.Time
func (d Date) YearDay() int {
	return d.ToTime().YearDay()
}

// IsZero reports whether d represents the zero date instant,
func (d Date) IsZero() bool {
	return d.Year == 0 && d.Month == time.Month(0) && d.Day == 0
}

// Unix returns d as a Unix time, the number of seconds elapsed
// Date.Unix() should calculate Unix time not including time.
func (d Date) Unix() int64 {
	return d.ToTime().Unix()
}
