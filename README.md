# xavi
[![GoDoc][1]][2]
[![GoCard][3]][4]
[![Build Status][5]][6]
[![codecov][7]][8]

[1]: https://godoc.org/github.com/atsushi-ishibashi/xavi?status.svg
[2]: https://godoc.org/github.com/atsushi-ishibashi/xavi
[3]: https://goreportcard.com/badge/github.com/atsushi-ishibashi/xavi
[4]: https://goreportcard.com/report/github.com/atsushi-ishibashi/xavi
[5]: https://travis-ci.org/atsushi-ishibashi/xavi.svg?branch=master
[6]: https://travis-ci.org/atsushi-ishibashi/xavi
[7]: https://codecov.io/gh/atsushi-ishibashi/xavi/branch/master/graph/badge.svg
[8]: https://codecov.io/gh/atsushi-ishibashi/xavi

xavi is a Go library for passing through elements from src struct to dst struct.

There are two ways to match elements
1. To use `xavi` struct tag. It requires those two type are the same.
2. Same name and type of field.

## Installation
```
$ go get github.com/atsushi-ishibashi/xavi
```

### Example
```
package main

import (
	"fmt"

	"github.com/atsushi-ishibashi/xavi"
)

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

	xavi.Pass(&dstHoge, hoge)

	fmt.Printf("%+v", dstHoge)
}

//Output
{Description:description ID:1 Name:name Struct:{Name:sub}}
```
