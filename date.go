package godate

import "time"

// ZeroDays represents 1, January, 1
const ZeroDays ElapsedDays = 0

// ElapsedDays represents elapsed date based on 1, January, 1
type ElapsedDays int64

// Date represents date under the proleptic Gregorian calendar used by ISO 8601.
// Zero value of Date is 1, January, 1.
// In v1.0.0, representation of BC (e.g. "-0912-01-01") is unsupported. (Please contribute!)
type Date struct {
	days ElapsedDays
}

// New returns Date from specified year, month, day.
func New(year int, month time.Month, day int) Date {
	t := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	return Date{toElapsedDays(t)}
}

// NewFromTime returns Date generated from time.Time's value
func NewFromTime(t time.Time) Date {
	return Date{toElapsedDays(t)}
}

// NewFromElapsedDays returns Date generated from ElapsedDate value
func NewFromElapsedDays(ed ElapsedDays) Date {
	return Date{ed}
}

// Today returns Date of today
func Today() Date {
	return Date{toElapsedDays(time.Now())}
}

// ToTime returns time.Time instance in UTC from d.
func (d Date) ToTime() time.Time {
	return toTime(d.days)
}

// Year returns year of date
func (d Date) Year() int {
	return d.ToTime().Year()
}

// Month returns month of date
func (d Date) Month() time.Month {
	return d.ToTime().Month()
}

// Day returns day of date
func (d Date) Day() int {
	return d.ToTime().Day()
}

// YearDay is wrapper of YearDay method of time.Time
func (d Date) YearDay() int {
	return d.ToTime().YearDay()
}

// IsZero reports whether d represents the zero date instant,
func (d Date) IsZero() bool {
	return d.days == ZeroDays
}

// Unix returns d as a Unix time, the number of seconds elapsed
// Date.Unix() should calculate Unix time not including time.
func (d Date) Unix() int64 {
	return d.ToTime().Unix()
}

const secondsPerDay int64 = 60 * 60 * 24

var zeroUnix = time.Time{}.Unix()

func toElapsedDays(t time.Time) ElapsedDays {
	elapsedSecs := t.Unix() - zeroUnix
	return ElapsedDays(elapsedSecs / secondsPerDay)
}

func toTime(days ElapsedDays) time.Time {
	return time.Unix(zeroUnix+(int64(days)*secondsPerDay), 0).UTC()
}
