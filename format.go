package godate

import "time"

// These are string format for Date
const (
	ANSIC    = "Jan _2 2006"
	RubyDate = "Jan 02 2006"
	RFC822   = "02 Jan 06"
	RFC3339  = "2006-01-02"
)

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

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a quoted string in RFC 3339 format.
func (d *Date) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	var err error
	*d, err = Parse(`"`+time.RFC3339+`"`, string(data))
	return err
}
