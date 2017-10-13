package gate

import "time"

type dateFormat string

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

// String returns formatted string
func (d Date) Format(format dateFormat) string {
	return d.t.Format(string(format))
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
