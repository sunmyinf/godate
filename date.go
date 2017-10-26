package godate

import "time"

// Date handles time.Time as date handler
type Date struct {
	t time.Time
}

// New returns Date from specified year, month, day and location
func New(year, month, day int, loc *time.Location) Date {
	return Date{time.Date(year, time.Month(month), day, 0, 0, 0, 0, loc)}
}

// Today returns Date of today
func Today() Date {
	return Date{
		time.Now(),
	}
}

// YearDay is wrapper of YearDay method of time.Time
func (d Date) YearDay() int {
	return d.t.YearDay()
}

// IsZero reports whether t represents the zero time instant,
// January 1, year 1, 00:00:00 UTC.
func (d Date) IsZero() bool {
	return d.t.IsZero()
}

// Year returns year of date
func (d Date) Year() int {
	return d.t.Year()
}

// Month returns month of date
func (d Date) Month() time.Month {
	return d.t.Month()
}

// Day returns day of date
func (d Date) Day() int {
	return d.t.Day()
}

// Unix returns t as a Unix time, the number of seconds elapsed
// Date.Unix() should calculate Unix time not including time.
func (d Date) Unix() int64 {
	return d.t.Round(time.Hour).Unix()
}
