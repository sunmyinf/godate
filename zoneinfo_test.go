package godate

import (
	"testing"
	"time"
)

const (
	format = time.RFC3339

	zoneUTC     = "UTC"
	zoneLocal   = "Local"
	zoneLA      = "America/Los_Angeles"
	zoneTokyo   = "Asia/Tokyo"
	zoneFakaofo = "Pacific/Fakaofo"

	zoneAbbrLA = "PDT"
)

func TestToday(t *testing.T) {
	today := Today()
	if today.t.IsZero() {
		t.Errorf("unexported t field of date is zero")
	}
}

func TestDate_In(t *testing.T) {
	tests := []struct {
		zoneName string
		want     string
	}{
		{
			zoneName: zoneUTC,
			want:     "2017-10-26T16:00:00Z",
		},
		{
			zoneName: zoneLA,
			want:     "2017-10-26T09:00:00-07:00",
		},
		{
			zoneName: zoneTokyo,
			want:     "2017-10-27T01:00:00+09:00",
		},
		{
			zoneName: zoneFakaofo,
			want:     "2017-10-27T05:00:00+13:00",
		},
	}

	for i, test := range tests {
		d, err := makeBaseDate()
		if err != nil {
			t.Error(err)
		}

		subject, err := d.In(test.zoneName)
		if err != nil {
			t.Error(err)
		}
		if subject.Format(format) != test.want {
			t.Errorf("expected %s, but got %s; test idx=%d", test.want, subject.Format(format), i)
		}
	}
}

func TestDate_UTC(t *testing.T) {
	d, err := makeBaseDate()
	if err != nil {
		t.Fatal(err)
	}
	d, err = d.In(zoneLA)
	if err != nil {
		t.Fatal(err)
	}

	subject := d.UTC()
	if zn, _ := subject.Zone(); zn != zoneUTC {
		t.Errorf("expected %s, but got %s", zoneUTC, zn)
	}
}

func TestDate_Local(t *testing.T) {
	d, err := makeBaseDate()
	if err != nil {
		t.Fatal(err)
	}
	subject := d.Local()

	if subject.Local().t.Location().String() != zoneLocal {
		t.Errorf("expected %s, but got %s",
			zoneLocal, subject.Local().t.Location().String())
	}
}

func TestDate_Location(t *testing.T) {
	d, err := makeBaseDate()
	if err != nil {
		t.Fatal(err)
	}
	if loc := d.Location(); loc.String() != zoneUTC {
		t.Errorf("expected %s, but got %s", zoneUTC, loc.String())
	}

	d, err = d.In(zoneLA)
	if err != nil {
		t.Fatal(err)
	}
	if loc := d.Location(); loc.String() != zoneLA {
		t.Errorf("expected %s, but got %s", zoneLA, loc.String())
	}

	d, err = d.In(zoneTokyo)
	if err != nil {
		t.Fatal(err)
	}
	if loc := d.Location(); loc.String() != zoneTokyo {
		t.Errorf("expected %s, but got %s", zoneTokyo, loc.String())
	}
}

func TestDate_Zone(t *testing.T) {
	d, err := makeBaseDate()
	if err != nil {
		t.Fatal(err)
	}
	zn, ofs := d.Zone()
	if zn != zoneUTC || ofs != 0 {
		t.Errorf(
			"expected name=%s, offset=%d; but got name=%s, offset=%d",
			zoneUTC, 0, zn, ofs)
	}

	d, err = d.In(zoneLA)
	if err != nil {
		t.Fatal(err)
	}
	zn, ofs = d.Zone()
	if zn != zoneAbbrLA || ofs != -25200 {
		t.Errorf(
			"expected name=%s, offset=%d; but got name=%s, offset=%d",
			zoneAbbrLA, -25200, zn, ofs)
	}
}
