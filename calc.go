package godate

import "time"

// Equal reports whether d and u represent the same time instant.
// This is an alias of == operation.
func (d Date) Equal(u Date) bool {
	return d.Year == u.Year &&
		d.Month == u.Month &&
		d.Day == u.Day
}

// After reports whether the time instant d is after u.
func (d Date) After(u Date) bool {
	if d.Equal(u) {
		return false
	}
	if d.Year > u.Year {
		return true
	} else if d.Year < u.Year {
		return false
	}
	return d.YearDay() > u.YearDay()
}

// Before reports whether the time instant d is before u.
func (d Date) Before(u Date) bool {
	if d.Equal(u) {
		return false
	}
	if d.Year < u.Year {
		return true
	} else if d.Year > u.Year {
		return false
	}
	return d.YearDay() < u.YearDay()
}

// Add year, month, day to d
func (d Date) Add(years, months, days int) Date {
	t := d.ToTime().AddDate(years, months, days)
	return Date{t.Year(), t.Month(), t.Day()}
}

const (
	// Day represents a day based on time.Duration
	Day time.Duration = 24 * time.Hour

	daysPerBasicYear time.Duration = 365 * Day
	daysPerLeapYear  time.Duration = 366 * Day
)

// Since returns the time elapsed since d.
// It is shorthand for godate.Today().Sub(d).
func Since(d Date) time.Duration {
	return Today().Sub(d)
}

// Sub returns the Duration resulted from d-u.
func (d Date) Sub(u Date) time.Duration {
	return d.ToTime().Sub(u.ToTime())
}
