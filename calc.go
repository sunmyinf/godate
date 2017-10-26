package godate

// Equal reports whether d and u represent the same time instant.
func (d Date) Equal(u Date) bool {
	return d.Year() == u.Year() &&
		d.Month() == u.Month() &&
		d.Day() == u.Day()
}

// After reports whether the time instant d is after u.
func (d Date) After(u Date) bool {
	return !(d.Equal(u) ||
		d.Year() < u.Year() ||
		d.Month() < u.Month() ||
		d.Day() < u.Day())
}

// Before reports whether the time instant d is before u.
func (d Date) Before(u Date) bool {
	return !(d.Equal(u) ||
		d.Year() > u.Year() ||
		d.Month() > u.Month() ||
		d.Day() > u.Day())
}

// Add year, month, day to d
func (d Date) Add(years, months, days int) Date {
	return Date{d.t.AddDate(years, months, days)}
}

// Since returns the time elapsed since d.
// It is shorthand for godate.Today().Sub(d).
func Since(d Date) int64 {
	return Today().Sub(d)
}

const (
	daysPerBasicYear int64 = 365
	daysPerLeapYear  int64 = 366
)

// Sub returns the days d-u.
func (d Date) Sub(u Date) (days int64) {
	dy, uy := d.Year(), u.Year()
	if dy == uy {
		return int64(d.YearDay() - u.YearDay())
	}

	if dy > uy {
		for y := uy; y <= dy; y++ {
			switch y {
			case uy:
				days += daysPer(y) - int64(u.YearDay())
			case dy:
				days += int64(d.YearDay())
			default:
				days += daysPer(y)
			}
		}
	} else {
		for y := uy; dy <= y; y-- {
			switch y {
			case uy:
				days -= int64(u.YearDay())
			case dy:
				days -= daysPer(y) - int64(d.YearDay())
			default:
				days -= daysPer(y)
			}
		}

	}
	return
}

func daysPer(year int) int64 {
	if isLeapYear(year) {
		return daysPerLeapYear
	}
	return daysPerBasicYear
}

// isLeap is implemented in time package
func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
