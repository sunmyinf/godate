package godate

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

// NullDate represents a date that may be null.
type NullDate struct {
	Date  Date
	Valid bool // Valid is true if Date is not NULL
}

// Scan implements the Scanner interface.
func (nd *NullDate) Scan(value interface{}) error {
	var err error
	switch x := value.(type) {
	case time.Time:
		nd.Date = NewFromTime(x)
	case nil:
		nd.Valid = false
		return nil
	default:
		err = fmt.Errorf("godate: cannot scan type %T into godate.NullDate: %v", value, value)
	}
	nd.Valid = err == nil
	return err
}

// Value implements the driver Valuer interface.
func (nd NullDate) Value() (driver.Value, error) {
	if !nd.Valid {
		return nil, nil
	}
	return nd.Date.String(), nil
}

// NewNullDate creates a new Date.
func NewNullDate(d Date, valid bool) NullDate {
	return NullDate{
		Date:  d,
		Valid: valid,
	}
}

// NullDateFrom creates a new NullDate that will always be valid.
func NullDateFrom(d Date) NullDate {
	return NewNullDate(d, true)
}

// NullDateFromPtr creates a new Date that will be null if d is nil.
func NullDateFromPtr(d *Date) NullDate {
	if d == nil {
		return NewNullDate(Date{}, false)
	}
	return NewNullDate(*d, true)
}

// MarshalJSON implements json.Marshaller.
// It will encode null if this date is null.
func (nd NullDate) MarshalJSON() ([]byte, error) {
	if !nd.Valid {
		return []byte("null"), nil
	}
	return nd.Date.MarshalJSON()
}

// UnmarshalJSON implements json.Unmarshaller.
func (nd *NullDate) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v.(type) {
	case string, time.Time:
		err = nd.Date.UnmarshalJSON(data)
	case nil:
		nd.Valid = false
		return nil
	default:
		err = fmt.Errorf("godate: cannot unmarshal %v into Go value of type godate.NullDate", reflect.TypeOf(v).Name())
	}
	nd.Valid = err == nil
	return err
}

// SetValid changes this Date's value and sets it to be non-null.
func (nd *NullDate) SetValid(v Date) {
	nd.Date = v
	nd.Valid = true
}

// Ptr returns a pointer to this Date's value, or a nil pointer if this Date is null.
func (nd NullDate) Ptr() *Date {
	if !nd.Valid {
		return nil
	}
	return &nd.Date
}

// String wraps Date.String(). If nd.Date is nil, return empty string.
func (nd NullDate) String() string {
	if !nd.Valid {
		return ""
	}
	return nd.Date.String()
}

// Format wraps Date.Format(). If nd.Date is nil, return empty string.
func (nd NullDate) Format(layout string) string {
	if !nd.Valid {
		return ""
	}
	return nd.Date.Format(layout)
}

// IsZero returns true for invalid Date.
func (nd NullDate) IsZero() bool {
	return !nd.Valid
}

// ValueOrZero returns the inner value if valid, otherwise zero.
func (nd NullDate) ValueOrZero() Date {
	if !nd.Valid {
		return Date{}
	}
	return nd.Date
}
