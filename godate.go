package godate

import "time"

type dateFormat string

// These are format for Date
const (
	ANSIC    dateFormat = "Jan _2 2006"
	RubyDate dateFormat = "Jan 02 2006"
	RFC822   dateFormat = "02 Jan 06"
	RFC3339  dateFormat = "2006-01-02"
)

// Date handles time.Time as date handler
type Date struct {
	t time.Time
}

// Today returns Date of today
func Today() Date {
	return Date{
		t: time.Now().Truncate(time.Hour * 24),
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
func (d Date) Format(format dateFormat) string {
	return d.t.Format(string(format))
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

// In returns Date in specified time zone
func (d Date) In(name string) (Date, error) {
	loc, err := time.LoadLocation(name)
	if err != nil {
		return Date{}, err
	}

	d.t = d.t.In(loc)
	return d, nil
}

// Local is Wrapper of Local method of time.Time
func (d Date) Local() Date {
	d.t = d.t.Local()
	return d
}

// UTC is Wrapper of UTC method of time.Time
func (d Date) UTC() Date {
	d.t = d.t.UTC()
	return d
}

// Zone is Wrapper of Zone method of time.Time
func (d Date) Zone() (name string, offset int) {
	return d.t.Zone()
}
