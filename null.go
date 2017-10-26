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
	Valid bool // Valid is true if Int64 is not NULL
}

// Scan implements the Scanner interface.
func (nd *NullDate) Scan(value interface{}) error {
	var err error
	switch x := value.(type) {
	case time.Time:
		nd.Date = Date{x}
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
	return nd.Date, nil
}

// NewNullDate creates a new Date.
func NewNullDate(d Date, valid bool) NullDate {
	return NullDate{
		Date:  d,
		Valid: valid,
	}
}

// NullDateFrom creates a new Time that will always be valid.
func NullDateFrom(d Date) NullDate {
	return NewNullDate(d, true)
}

// MarshalJSON implements json.Marshaler.
// It will encode null if this date is null.
func (nd NullDate) MarshalJSON() ([]byte, error) {
	if !nd.Valid {
		return []byte("null"), nil
	}
	return nd.Date.MarshalJSON()
}

// UnmarshalJSON implements json.Unmarshaler.
func (nd *NullDate) UnmarshalJSON(data []byte) error {
	var err error
	var v interface{}
	if err = json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v.(type) {
	case string:
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
