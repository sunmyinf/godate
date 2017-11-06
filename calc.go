package godate

import "time"

// Equal reports whether d and u represent the same time instant.
// This is an alias of == operation.
func (d Date) Equal(u Date) bool {
	return d == u
}

// After reports whether the time instant d is after u.
func (d Date) After(u Date) bool {
	if d.Year != u.Year {
		return d.Year > u.Year
	}
	if d.Month != u.Month {
		return d.Month > u.Month
	}
	return d.Day > u.Day
}

// Before reports whether the time instant d is before u.
func (d Date) Before(u Date) bool {
	if d.Year != u.Year {
		return d.Year < u.Year
	}
	if d.Month != u.Month {
		return d.Month < u.Month
	}
	return d.Day < u.Day
}

// Add year, month, day to d
func (d Date) Add(years, months, days int) Date {
	t := d.ToTime().AddDate(years, months, days)
	return Date{t.Year(), t.Month(), t.Day()}
}

// Day represents a day based on time.Duration
const Day time.Duration = 24 * time.Hour

// Since returns the time elapsed since d.
// It is shorthand for godate.Today().Sub(d).
func Since(d Date) time.Duration {
	return Today().Sub(d)
}

// Sub returns the Duration resulted from d-u.
func (d Date) Sub(u Date) time.Duration {
	return d.ToTime().Sub(u.ToTime())
}
