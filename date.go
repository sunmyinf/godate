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
