package godate

import "time"

// In returns Date in specified time zone
func (d Date) In(name string) (Date, error) {
	loc, err := time.LoadLocation(name)
	if err != nil {
		return Date{}, err
	}

	d.t = d.t.In(loc)
	return d, nil
}

// Local is Wrapper of Local method of time.Time
func (d Date) Local() Date {
	d.t = d.t.Local()
	return d
}

// Location returns the time zone information associated with Date.t.
func (d Date) Location() *time.Location {
	return d.t.Location()
}

// UTC is Wrapper of UTC method of time.Time
func (d Date) UTC() Date {
	d.t = d.t.UTC()
	return d
}

// Zone is Wrapper of Zone method of time.Time
func (d Date) Zone() (name string, offset int) {
	return d.t.Zone()
}
