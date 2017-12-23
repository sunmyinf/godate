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
	t, err := time.Parse(layout, value)
	if err != nil {
		return Date{}, err
	}
	return NewFromTime(t), nil
}

// Format is wrapper of Format method of time.Time
// When IsZero is true, returns an empty string
func (d Date) Format(format string) string {
	return d.ToTime().Format(format)
}

// String returns string of RFC3339 date format
func (d Date) String() string {
	return d.Format(RFC3339)
}

// AppendFormat is wrapper of AppendFormat method of time.Time
func (d Date) AppendFormat(b []byte, layout string) []byte {
	return d.ToTime().AppendFormat(b, layout)
}

// MarshalJSON implements the json.Marshaller interface.
// The date is a quoted string in RFC 3339 format.
func (d Date) MarshalJSON() ([]byte, error) {
	if y := d.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		return nil, errors.New("Date.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(RFC3339)+2)
	b = append(b, '"')
	b = d.AppendFormat(b, RFC3339)
	b = append(b, '"')
	return b, nil
}

// UnmarshalJSON implements the json.Unmarshaller interface.
// The date is expected to be a quoted string in RFC 3339 format.
func (d *Date) UnmarshalJSON(data []byte) error {
	var err error
	*d, err = Parse(`"`+RFC3339+`"`, string(data))
	return err
}

// Value implements the driver Valuer interface.
func (d Date) Value() (driver.Value, error) {
	return d.String(), nil
}

// Scan implements the sql.Scanner interface.
func (d *Date) Scan(value interface{}) error {
	var err error
	switch x := value.(type) {
	case string:
		*d, err = Parse(RFC3339, x)
		if err != nil {
			return err
		}
		return nil
	case time.Time:
		*d = NewFromTime(x)
		return nil
	default:
		return fmt.Errorf("godate: cannot scan type %T into godate.Date: %v", value, value)
	}
}
