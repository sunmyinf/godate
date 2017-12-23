package godate

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	d, err := Parse(RFC3339, "2017-10-20")
	if err != nil {
		t.Errorf("expected error not raised, but got=%v", err)
	}
	if d.String() != "2017-10-20" {
		t.Errorf("expected String() returns '2017-10-20', but got %s", d.String())
	}

	d, err = Parse("20060102", "20171020")
	if err != nil {
		t.Errorf("expected error not raised, but got=%v", err)
	}
	if d.String() != "2017-10-20" {
		t.Errorf("expected String() returns '2017-10-20', but got %s", d.String())
	}

	_, err = Parse("20161020", "2017-10-20")
	if err == nil {
		t.Errorf("expected error raised, bug got %v", err)
	}
}

func TestDate_Format(t *testing.T) {
	d := New(2017, 10, 27)
	if d.Format(RFC3339) != "2017-10-27" {
		t.Errorf("expected formatted godate.RFC3339, but got %s", d.Format(RFC3339))
	}
	if d.Format("2006/01/02") != "2017/10/27" {
		t.Errorf("expected formatted to '2006/01/02', but got %s", d.Format("2006/01/02"))
	}

	d = Date{}
	if d.Format(RFC3339) != "0001-01-01" {
		t.Errorf("expected zero date is formatted as empty, but got %s", d.Format(RFC3339))
	}
}

func TestDate_String(t *testing.T) {
	d := New(2017, 10, 27)
	if d.String() != "2017-10-27" {
		t.Errorf("expected formatted godate.RFC3339, but got %s", d.String())
	}

	d = Date{}
	if d.Format(RFC3339) != "0001-01-01" {
		t.Errorf("expected zero date is formatted as empty, but got %s", d.Format(RFC3339))
	}
}

func TestDate_MarshalJSON(t *testing.T) {
	d := New(2017, 10, 27)
	if buf, err := d.MarshalJSON(); err != nil {
		t.Errorf("expected to be marshaled, but got error %v", err)
	} else if string(buf) != "\"2017-10-27\"" {
		t.Errorf("expected formatted json string, but got %s", buf)
	}

	d = Date{}
	if buf, err := d.MarshalJSON(); err != nil {
		t.Errorf("expected to be marshaled, but got error %v", err)
	} else if bytes.Equal(buf, []byte("0001-01-01")) {
		fmt.Println(string(buf))
		fmt.Println(string(buf) != "0001-01-01")
		t.Errorf("expected formatted json string, but got %s", buf)
	}
}

func TestDate_UnmarshalJSON(t *testing.T) {
	d := Date{}
	if err := d.UnmarshalJSON([]byte("\"2017-10-27\"")); err != nil {
		t.Errorf("expected to be unmarshaled, but got error %v", err)
	} else if d != New(2017, 10, 27) {
		t.Errorf("expected parsed date, but got %s", d)
	}

	d = New(2017, 10, 28)
	if err := d.UnmarshalJSON([]byte("null")); err == nil {
		t.Errorf("expected to be raised error, but not got error")
	}
}

func TestDate_Value(t *testing.T) {
	d := New(2017, 10, 27)
	if v, err := d.Value(); err != nil {
		t.Errorf("expected to get sql value, but got error %v", err)
	} else if v != "2017-10-27" {
		t.Errorf("expected formatted string, but got %s", v)
	}

	d = Date{}
	if v, err := d.Value(); err != nil {
		t.Errorf("expected to get sql value, but got error %v", err)
	} else if v != "0001-01-01" {
		t.Errorf("expected nil, but got %s", v)
	}
}

func TestDate_Scan(t *testing.T) {
	d := Date{}
	if err := d.Scan("2017-10-27"); err != nil {
		t.Errorf("expected to be scanned, but got error %v", err)
	} else if d != New(2017, 10, 27) {
		t.Errorf("expected parsed date, but got %s", d)
	}

	d = Date{}
	if err := d.Scan(time.Date(2017, time.October, 27, 0, 0, 0, 0, time.UTC)); err != nil {
		t.Errorf("expected to be scanned, but got error %v", err)
	} else if d != New(2017, 10, 27) {
		t.Errorf("expected parsed date, but got %s", d)
	}

	d = New(2017, 10, 28)
	if err := d.Scan(nil); err == nil {
		t.Errorf("expected not to be scanned, but scanned")
	}
	if err := d.Scan("2017/10/28"); err == nil {
		t.Errorf("expected not to be scanned, but scanned")
	}
}
