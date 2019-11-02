# xavi
[![GoDoc][1]][2]

[1]: https://godoc.org/github.com/atsushi-ishibashi/xavi?status.svg
[2]: https://godoc.org/github.com/atsushi-ishibashi/xavi

xavi is a Go library for passing through elements from src struct to dst struct.

There are two ways to match elements
1. To use `xavi` struct tag. It requires those two type are the same.
2. Field name and type is the same.

## Installation
```
$ go get github.com/atsushi-ishibashi/xavi
```

### Example
```
type Hoge struct {
    Description string

	HogeID     int64   `xavi:"id"`
	HogeName   string  `xavi:"name"`
	HogeStruct SubHoge `xavi:"subHoge"`
}

type SubHoge struct {
	Name string
}

type DstHoge struct {
    Description string

	ID     int64   `xavi:"id"`
	Name   string  `xavi:"name"`
	Struct SubHoge `xavi:"subHoge"`
}

func main() {
    hoge := Hoge{
		HogeID:   1,
		HogeName: "name",
		HogeStruct: SubHoge{
			Name: "sub",
		},
        Description: "description",
	}
    var dstHoge DstHoge
    
    err := xavi.Pass(&dstHoge, hoge)
}
```