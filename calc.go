package godate

// Since returns the time elapsed since d.
// It is shorthand for godate.Today().Sub(d).
func Since(d Date) int64 {
	return Today().Sub(d)
}

// Equal reports whether d and u represent the same time instant.
func (d Date) Equal(u Date) bool {
	return d.Year() == u.Year() &&
		d.Month() == u.Month() &&
		d.Day() == u.Day()
}

// After reports whether the time instant d is after u.
func (d Date) After(u Date) bool {
	return d.Year() >= u.Year() &&
		d.Month() >= u.Month() &&
		d.Day() > u.Day()
}

// Before reports whether the time instant d is before u.
func (d Date) Before(u Date) bool {
	return d.Year() <= u.Year() &&
		d.Month() <= u.Month() &&
		d.Day() < u.Day()
}

// Add year, month, day to d
func (d Date) Add(years, months, days int) Date {
	return Date{d.t.AddDate(years, months, days)}
}

// Sub returns the days d-u.
func (d Date) Sub(u Date) int64 {
	dyear := d.Year()
	dateyear := u.Year()
	diffYearDay := int64(u.YearDay() - d.YearDay())

	if dyear == dateyear {
		return diffYearDay
	}

	var diff int64
	if dyear < dateyear {
		for y := dateyear; dyear <= y; y-- {
			if y == dyear {
				diff += diffYearDay
			} else {
				diff += int64(baseDaysInYear(y))
			}
		}
	} else {
		for y := dateyear; y <= dyear; y++ {
			if y == dyear {
				diff += diffYearDay
			} else {
				diff += int64(baseDaysInYear(y))
			}
		}
	}
	return diff
}

const (
	daysInBasicYear int = 365
	daysInLeapYear  int = 366
)

func baseDaysInYear(year int) int {
	if isLeapYear(year) {
		return daysInLeapYear
	}
	return daysInBasicYear
}

// isLeap is implemented in time package
func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
