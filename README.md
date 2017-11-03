# godate
[![Build Status](https://travis-ci.org/sunmyinf/godate.svg?branch=travis_ci)](https://travis-ci.org/sunmyinf/godate)
[![Go Report Card](https://goreportcard.com/badge/github.com/sunmyinf/godate)](https://goreportcard.com/report/github.com/sunmyinf/godate)
[![GoDoc](https://godoc.org/github.com/sunmyinf/godate?status.svg)](https://godoc.org/github.com/sunmyinf/godate)  

## Overview
godate has `Date` and `NullDate` type to handle date in go.   
`Date` type is so simple so that we  have not to be concerned with TimeZone or location. Also, `Date` is compatible with `time.Time`.  
For example, `Date.ToTime()` returns time.Time instance of UTC location that have only Year, Month, Day.

## Getting Started     

### Basic `Date` Usage
```go
d, _ := godate.Parse(godate.RFC3339, "2017-10-13")
fmt.Println(d) // => "2017-10-13"

// Sub returns days resulted from d - u
u := d.Add(1,0,1)
fmt.Println(d.Sub(u)) // => 366
 
// Format
fmt.Println(u.Format(godate.RubyDate)) // => "Oct 13 2017"

// ToTime returns time.Time instance from Date fields value.
// the returned Time instance is in UTC time zone.
t := d.ToTime()
fmt.Println(t.Format(time.RFC3339)) // => "2017-10-13T00:00:00Z" 
```
Also, implemented formatting methods like `UnmarshalJSON`, `MarshalJSON`, `Scan`, `Value` and so on.

## Contribution

1. Fork ([https://github.com/sunmyinf/godate](https://github.com/sunmyinf/godate))
2. Create a feature branch
3. Commit your changes
4. Rebase your local changes against the master branch
5. Run test suite with the `go test ./...` command and confirm that it passes
6. Run `gofmt -s` or `goimports -s`
7. Create new Pull Request

## License
godate is released under the [MIT License](https://opensource.org/licenses/MIT).
