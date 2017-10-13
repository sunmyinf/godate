# gate
`gate` is package wrapping time.Time to handle date easily in go.

## Get Started
### Installation
```bash
$ go get -u github.com/sunmyinf/gate
```

### Usage
```go
package main

import (
    "fmt"

    "github.com/sunmyinf/gate"
)

func main() {
    date, _ := gate.Parse("2006,01,02", "2017,10,13")
    fmt.Println(date.Format(gate.RFC3339)) // "2017-10-13"
}
```
