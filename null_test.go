package godate

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestNewNullDate(t *testing.T) {
	nd := NewNullDate(New(2017, 10, 27), true)
	if !nd.Valid {
		t.Errorf("expected valid, but invalid")
	}
	if nd.Date.String() != "2017-10-27" {
		t.Errorf("expected '2017-10-27', but got %s", nd.Date.String())
	}

	nd = NewNullDate(Date{}, false)
	if nd.Valid {
		t.Errorf("expected invalid, but valid")
	}
	if !nd.Date.IsZero() {
		t.Errorf("expected date is zero, but not zero")
	}
}

func TestNullDateFrom(t *testing.T) {
	nd := NullDateFrom(New(2017, 10, 27))
	if !nd.Valid {
		t.Errorf("expected valid, but invalid")
	}
	if nd.Date.String() != "2017-10-27" {
		t.Errorf("expected '2017-10-27' returned, but got %s", nd.Date.String())
	}
}

func TestNullDateFromPtr(t *testing.T) {
	d := New(2017, 10, 27)
	nd := NullDateFromPtr(&d)
	if !nd.Valid {
		t.Errorf("expected valid, but invalid")
	}
	if nd.Date.String() != "2017-10-27" {
		t.Errorf("expected '2017-10-27' returned, but got %s", nd.Date.String())
	}

	nd = NullDateFromPtr(nil)
	if nd.Valid {
		t.Errorf("expected invalid, but valid")
	}
	if !nd.Date.IsZero() {
		t.Errorf("expected date is zero, but not zero")
	}
}

func TestNullDate_SetValid(t *testing.T) {
	nd := NullDateFromPtr(nil)
	if nd.Valid {
		t.Fatal("unexpected nd.Valid is true")
	}

	nd.SetValid(New(2017, 10, 27))
	if !nd.Valid {
		t.Errorf("expected valid, but invalid")
	}
	if nd.Date.String() != "2017-10-27" {
		t.Errorf("expected '2017-10-27' returned, but got %s", nd.Date.String())
	}
}

func TestNullDate_Ptr(t *testing.T) {
	d := New(2017, 10, 27)
	nd := NullDateFrom(d)
	dp := nd.Ptr()

	v := reflect.ValueOf(dp)
	if v.Kind() != reflect.Ptr {
		t.Errorf("expected Ptr returns pointer, but got %v", v.Kind())
	}
}

func TestNullDate_String(t *testing.T) {
	nd := NullDateFromPtr(nil)
	if nd.String() != "" {
		t.Errorf("expected receive empty string, but got %s", nd.String())
	}

	d := New(2017, 10, 27)
	nd = NullDateFrom(d)
	if nd.String() != "2017-10-27" {
		t.Errorf("expected receive '2017-10-27', but got %s", nd.String())
	}
}

func TestNullDate_Format(t *testing.T) {
	nd := NullDateFromPtr(nil)
	if nd.Format(RFC3339) != "" {
		t.Errorf("expected receive empty string, but got %s", nd.String())
	}

	d := New(2017, 10, 27)
	nd = NullDateFrom(d)
	if nd.Format("2006/01/02") != "2017/10/27" {
		t.Errorf("expected receive '2017/10/27', but got %s", nd.String())
	}
}

func TestNullDate_IsZero(t *testing.T) {
	nd := NewNullDate(Date{}, true)
	if nd.IsZero() {
		t.Errorf("expected NullDate.IsZero is false but got true")
	}

	nd.Valid = false
	if !nd.IsZero() {
		t.Errorf("expected NullDate.IsZero is true but got false")
	}
}

func TestNullDate_ValueOrZero(t *testing.T) {
	d := New(2017, 12, 15)
	nd := NewNullDate(d, true)
	if nd.ValueOrZero() != d {
		t.Errorf("expected ValueOrZero() is equal to d, but not equal. ValueOrZero=%v, d=%v", nd.ValueOrZero(), d)
	}

	nd.Valid = false
	if !nd.ValueOrZero().IsZero() {
		t.Errorf("expected ValueOrZero is zero. but got zero")
	}
}

func TestNullDate_Value(t *testing.T) {
	d := New(2017, 12, 16)
	nd := NewNullDate(d, true)

	value, err := nd.Value()
	if err != nil {
		t.Error(err)
	}
	if value != d.String() {
		t.Errorf("unexpected Value's value. actual=%v, expected=%v", value, d)
	}

	nd.Valid = false
	value, err = nd.Value()
	if err != nil {
		t.Error(err)
	}
	if value != nil {
		t.Errorf("unexpected Value's value. actual=%v, expected=nil", value)
	}
}

func TestNullDate_Scan(t *testing.T) {
	tests := []struct {
		value         interface{}
		expected      NullDate
		errorExpected bool
	}{
		{
			value:         "2017-12-10",
			expected:      NullDateFrom(New(2017, 12, 10)),
			errorExpected: false,
		},
		{
			value:         time.Date(2017, 12, 16, 0, 0, 0, 0, time.UTC),
			expected:      NullDateFrom(New(2017, 12, 16)),
			errorExpected: false,
		},
		{
			value:         nil,
			expected:      NewNullDate(Date{}, false),
			errorExpected: false,
		},
		{
			value:         6,
			errorExpected: true,
		},
	}

	for i, te := range tests {
		actual := NullDate{}
		err := actual.Scan(te.value)
		if err != nil {
			if !te.errorExpected {
				t.Errorf("got unexpected error. err=%v", err)
			}
			continue
		} else {
			if te.errorExpected {
				t.Errorf("expected error, but not got.idx=%d", i)
			}
		}

		if te.expected != actual {
			t.Errorf("unexpected NullDate. idx=%d, actual=%v, expected=%v", i, actual, te.expected)
		}
	}
}

func TestNullDate_MarshalJSON(t *testing.T) {
	d := New(2017, 12, 15)
	nd := NullDateFrom(d)

	b, err := nd.MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(b, []byte("\""+d.String()+"\"")) {
		t.Errorf("unexpected MarshalJSON value. actual=%v, expected=%v", b, d)
	}

	nd.Valid = false
	b, err = nd.MarshalJSON()
	if !bytes.Equal(b, []byte("null")) {
		t.Errorf("unexpected MarshalJSON value. actual=%v, expected=null", b)
	}
}

func TestNullDate_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		data          []byte
		expected      NullDate
		errorExpected bool
	}{
		{
			data:          []byte("\"2017-12-10\""),
			expected:      NullDateFrom(New(2017, 12, 10)),
			errorExpected: false,
		},
		{
			data:          []byte("null"),
			expected:      NewNullDate(Date{}, false),
			errorExpected: false,
		},
		{
			data:          []byte("6"),
			errorExpected: true,
		},
	}

	for i, te := range tests {
		actual := NullDate{}
		err := actual.UnmarshalJSON(te.data)
		if err != nil {
			if !te.errorExpected {
				t.Errorf("got unexpected error. err=%v", err)
			}
			continue
		} else {
			if te.errorExpected {
				t.Errorf("expected error, but not got.idx=%d", i)
			}
		}

		if te.expected != actual {
			t.Errorf("unexpected NullDate. idx=%d, actual=%v, expected=%v", i, actual, te.expected)
		}
	}
}
