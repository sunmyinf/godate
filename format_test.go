package godate

import (
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

	d, err = Parse("20060102Z", "20171020JST")
	if err != nil {
		t.Errorf("expected error not raised, but got=%v", err)
	}
	if d.String() != "2017-10-20" {
		t.Errorf("expected String() returns '2017-10-20', but got %s", d.String())
	}
	if zn, _ := d.Zone(); zn != "JST" {
		t.Errorf("expected zone name 'JST', but got %s", zn)
	}

	_, err = Parse("20161020", "2017-10-20")
	if err == nil {
		t.Errorf("expected error raised, bug got %v", err)
	}
}

func TestDate_Format(t *testing.T) {
	d := New(2017, 10, 27, time.UTC)
	if d.Format(RFC3339) != "2017-10-27" {
		t.Errorf("expected formatted godate.RFC3339, but got %s", d.Format(RFC3339))
	}
	if d.Format("2006/01/02") != "2017/10/27" {
		t.Errorf("expected formatted to '2006/01/02', but got %s", d.Format("2006/01/02"))
	}
}

func TestDate_String(t *testing.T) {
	d := New(2017, 10, 27, time.UTC)
	if d.String() != "2017-10-27" {
		t.Errorf("expected formatted godate.RFC3339, but got %s", d.String())
	}
}
