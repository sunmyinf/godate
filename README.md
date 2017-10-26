# godate
[![GoDoc](https://godoc.org/github.com/sunmyinf/godate?status.svg)](https://godoc.org/github.com/sunmyinf/godate)
[![Go Report Card](https://goreportcard.com/badge/github.com/sunmyinf/godate)](https://goreportcard.com/report/github.com/sunmyinf/godate)  
`godate` is package wrapping time.Time to handle date easily in go.

## Getting Started
godate has `Date` and `NullDate` type, wrapping `time.Time` type, to handle date in go.   

### Basic `Date` Usage
```go
package main

import (
    "fmt"
    "time"
    "github.com/sunmyinf/godate"
)

func main() {
    d, _ := godate.Parse(godate.RFC3339, "2017-10-13")
    fmt.Println(d) // => "2017-10-13"
    
    // Sub returns days resulted from d - u
    u := d.Add(1,0,1)
    fmt.Println(d.Sub(u)) // => 366
     
    // Format
    fmt.Println(u.Format(godate.RubyDate)) // => "Oct 13 2017"
}
```
Also, implemented `UnmarshalJSON`, `MarshalJSON`, `Scan`, `Value` and so on.

