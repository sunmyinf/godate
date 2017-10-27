package godate

import "time"

// Equal reports whether d and u represent the same time instant.
func (d Date) Equal(u Date) bool {
	return d.Year() == u.Year() &&
		d.Month() == u.Month() &&
		d.Day() == u.Day()
}

// After reports whether the time instant d is after u.
func (d Date) After(u Date) bool {
	if d.Equal(u) {
		return false
	}
	if d.Year() > u.Year() {
		return true
	} else if d.Year() < u.Year() {
		return false
	}
	return d.YearDay() > u.YearDay()
}

// Before reports whether the time instant d is before u.
func (d Date) Before(u Date) bool {
	if d.Equal(u) {
		return false
	}
	if d.Year() < u.Year() {
		return true
	} else if d.Year() > u.Year() {
		return false
	}
	return d.YearDay() < u.YearDay()
}

// Add year, month, day to d
func (d Date) Add(years, months, days int) Date {
	return Date{d.t.AddDate(years, months, days)}
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

// Sub returns the days d-u.
func (d Date) Sub(u Date) time.Duration {
	dy, uy := d.Year(), u.Year()
	if dy == uy {
		return time.Duration(d.YearDay()-u.YearDay()) * Day
	}

	var dd time.Duration
	if dy > uy {
		for y := uy; y <= dy; y++ {
			switch y {
			case uy:
				dd += daysPer(y) - time.Duration(u.YearDay())*Day
			case dy:
				dd += time.Duration(d.YearDay()) * Day
			default:
				dd += daysPer(y)
			}
		}
	} else {
		for y := uy; dy <= y; y-- {
			switch y {
			case uy:
				dd -= time.Duration(u.YearDay()) * Day
			case dy:
				dd -= daysPer(y) - time.Duration(d.YearDay())*Day
			default:
				dd -= daysPer(y)
			}
		}

	}
	return dd
}

func daysPer(year int) time.Duration {
	if isLeapYear(year) {
		return daysPerLeapYear
	}
	return daysPerBasicYear
}

// isLeap is implemented in time package
func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
