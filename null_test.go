package godate

import (
	"reflect"
	"testing"
	"time"
)

func TestNewNullDate(t *testing.T) {
	nd := NewNullDate(New(2017, 10, 27, time.UTC), true)
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
	nd := NullDateFrom(New(2017, 10, 27, time.UTC))
	if !nd.Valid {
		t.Errorf("expected valid, but invalid")
	}
	if nd.Date.String() != "2017-10-27" {
		t.Errorf("expected '2017-10-27' returned, but got %s", nd.Date.String())
	}
}

func TestNullDateFromPtr(t *testing.T) {
	d := New(2017, 10, 27, time.UTC)
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

	nd.SetValid(New(2017, 10, 27, time.UTC))
	if !nd.Valid {
		t.Errorf("expected valid, but invalid")
	}
	if nd.Date.String() != "2017-10-27" {
		t.Errorf("expected '2017-10-27' returned, but got %s", nd.Date.String())
	}
}

func TestNullDate_Ptr(t *testing.T) {
	d := New(2017, 10, 27, time.UTC)
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

	d := New(2017, 10, 27, time.UTC)
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

	d := New(2017, 10, 27, time.UTC)
	nd = NullDateFrom(d)
	if nd.Format("2006/01/02") != "2017/10/27" {
		t.Errorf("expected receive '2017/10/27', but got %s", nd.String())
	}
}
