# godate
[![GoDoc](https://godoc.org/github.com/sunmyinf/godate?status.svg)](https://godoc.org/github.com/sunmyinf/godate)  
`godate` is package wrapping time.Time to handle date easily in go.

## Get Started
### Installation
```bash
$ go get -u github.com/sunmyinf/godate
```

### Usage
```go
package main

import (
    "fmt"
    "github.com/sunmyinf/godate"
)

func main() {
    date, _ := godate.Parse("2006,01,02", "2017,10,13")
    fmt.Println(date.Format(godate.RFC3339)) // "2017-10-13"
}
```
