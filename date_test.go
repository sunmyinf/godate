package godate

import "testing"

func TestToday(t *testing.T) {
	today := Today()
	if today.t.IsZero() {
		t.Errorf("unexported t field of date is zero")
	}
}

func TestParse(t *testing.T) {

}

func TestDate_Format(t *testing.T) {

}

func TestDate_String(t *testing.T) {

}

func TestDate_YearDay(t *testing.T) {

}

func TestDate_IsZero(t *testing.T) {

}

func makeFixedDate() (Date, error) {
	d, err := Parse("2006-01-02T15:04:05Z", "2017-10-26T16:00:00Z")
	if err != nil {
		return Date{}, err
	}
	return d, err
}
