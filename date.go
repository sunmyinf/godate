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
		t: time.Now(),
	}
}

// Parse returns Date from time.Parse value
func Parse(layout, value string) (Date, error) {
	p, err := time.Parse(layout, value)
	if err != nil {
		return Date{}, err
	}

	return Date{t: p}, nil
}

// Format is wrapper of Format method of time.Time
func (d Date) Format(format string) string {
	return d.t.Format(format)
}

// String returns string of RFC3339 date format
func (d Date) String() string {
	return d.Format(RFC3339)
}

// Equal returns bool resulted from comparing two date object
func (d Date) Equal(date Date) bool {
	year, month, day := d.t.Date()
	xyear, xmonth, xday := date.t.Date()

	return year == xyear && month == xmonth && day == xday
}

// After reports whether the time instant d is after date.
func (d Date) After(date Date) bool {
	year, month, day := d.t.Date()
	xyear, xmonth, xday := date.t.Date()

	return year >= xyear && month >= xmonth && day > xday
}

// Before reports whether the time instant d is before date.
func (d Date) Before(date Date) bool {
	year, month, day := d.t.Date()
	xyear, xmonth, xday := date.t.Date()

	return year <= xyear && month <= xmonth && day < xday
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
