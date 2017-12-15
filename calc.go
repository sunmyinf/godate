package godate

// Equal reports whether d and u represent the same time instant.
// This is an alias of == operator.
func (d Date) Equal(u Date) bool {
	return d == u
}

// After reports whether the time instant d is after u.
func (d Date) After(u Date) bool {
	return d.days > u.days
}

// Before reports whether the time instant d is before u.
func (d Date) Before(u Date) bool {
	return d.days < u.days
}

// Add year, month, day to d
func (d Date) Add(years, months, days int) Date {
	t := d.ToTime().AddDate(years, months, days)
	return Date{toElapsedDays(t)}
}

// Since returns the days elapsed since d.
// It is shorthand for godate.Today().Sub(d).
func Since(d Date) ElapsedDays {
	return Today().Sub(d)
}

// Sub returns the Duration resulted from d-u.
func (d Date) Sub(u Date) ElapsedDays {
	return d.days - u.days
}
