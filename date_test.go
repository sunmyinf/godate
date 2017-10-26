package godate

func makeBaseDate() (Date, error) {
	d, err := Parse("2006-01-02T15:04:05Z", "2017-10-26T16:00:00Z")
	if err != nil {
		return Date{}, err
	}
	return d, err
}
