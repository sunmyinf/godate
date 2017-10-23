package godate

import "time"

// These are string format for Date
const (
	ANSIC    = "Jan _2 2006"
	RubyDate = "Jan 02 2006"
	RFC822   = "02 Jan 06"
	RFC3339  = "2006-01-02"
)

// Date handles time.Time as date handler
type Date struct {
	t time.Time
}

// Today returns Date of today
func Today() Date {
	return Date{
		time.Now(),
	}
}

// Parse returns Date from time.Parse value
func Parse(layout, value string) (Date, error) {
	p, err := time.Parse(layout, value)
	if err != nil {
		return Date{}, err
	}

	return Date{p}, nil
}

// Format is wrapper of Format method of time.Time
func (d Date) Format(format string) string {
	return d.t.Format(format)
}

// String returns string of RFC3339 date format
func (d Date) String() string {
	return d.Format(RFC3339)
}

// Equal reports whether d and date represent the same date instant.
func (d Date) Equal(date Date) bool {
	return d.t.Year() == date.t.Year() &&
		d.t.Month() == date.t.Month() &&
		d.t.Day() == date.t.Day()
}

// After reports whether the date instant d is after date.
func (d Date) After(date Date) bool {
	return d.t.Year() >= date.t.Year() &&
		d.t.Month() >= date.t.Month() &&
		d.t.Day() > date.t.Day()
}

// Before reports whether the date instant d is before date.
func (d Date) Before(date Date) bool {
	return d.t.Year() <= date.t.Year() &&
		d.t.Month() <= date.t.Month() &&
		d.t.Day() < date.t.Day()
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

// Add year, month, day to d
func (d Date) Add(years, months, days int) Date {
	return Date{d.t.AddDate(years, months, days)}
}

func (d Date) Diff(date Date) int64 {
	dyear := d.Year()
	dateyear := date.Year()
	diffYearDay := int64(date.YearDay() - d.YearDay())

	if dyear == dateyear {
		return diffYearDay
	}

	var diff int64
	if dyear < dateyear {
		for y := dateyear; dyear <= y; y-- {
			if y == dyear {
				diff += diffYearDay
			} else {
				diff += int64(baseDaysInYear(y))
			}
		}
	} else {
		for y := dateyear; y <= dyear; y++ {
			if y == dyear {
				diff += diffYearDay
			} else {
				diff += int64(baseDaysInYear(y))
			}
		}
	}
	return diff
}

const (
	daysInBasicYear int = 365
	daysInLeapYear  int = 366
)

func baseDaysInYear(year int) int {
	if isLeapYear(year) {
		return daysInLeapYear
	}
	return daysInBasicYear
}

// isLeap is implemented in time package
func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
