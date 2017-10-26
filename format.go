package godate

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
)

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

// MarshalJSON implements the json.Marshaler interface.
// The time is a quoted string in RFC 3339 format.
func (d Date) MarshalJSON() ([]byte, error) {
	if y := d.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(RFC3339)+2)
	b = append(b, '"')
	b = append(b, byte(d.Format(RFC3339)))
	b = append(b, '"')
	return b, nil
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

// Value implements the driver Valuer interface.
func (d Date) Value() (driver.Value, error) { return d.String(), nil }

// Scan implements the sql.Scanner interface.
func (d *Date) Scan(value interface{}) error {
	var err error
	switch x := value.(type) {
	case string:
		*d, err = Parse(RFC3339, x)
		if err != nil {
			return err
		}
	case []byte:
		*d, err = Parse(RFC3339, string(x))
		if err != nil {
			return err
		}
	case time.Time:
		*d = Date{x}
		return nil
	default:
		return fmt.Errorf("godate: cannot scan type %T into godate.Date: %v", value, value)
	}
}
