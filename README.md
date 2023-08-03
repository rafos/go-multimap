[![GoDoc](https://godoc.org/github.com/rafos/go-multimap?status.svg)](https://godoc.org/github.com/rafos/go-multimap) [![Go Report Card](https://goreportcard.com/badge/github.com/rafos/go-multimap)](https://goreportcard.com/report/github.com/rafos/go-multimap) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/rafos/go-multimap/blob/main/LICENSE)

# Go-Multimap (Generic)

This is my very first attempt at implementation the missing `multimap` data structure for the [Go](https://www.golang.org/project/) language.
It's heavily inspired by Jason's Wangsadinata [go-multimap](https://github.com/jwangsadinata/go-multimap).

The primary impetus, on the other hand, is that I come from programming in the Java language, where there is a [Guava](https://github.com/google/guava) library with a multimap implementation.

References:
[Wikipedia](https://en.wikipedia.org/wiki/Multimap),
[Guava](https://google.github.io/guava/releases/19.0/api/docs/com/google/common/collect/Multimap.html)

## Installation ##

Install the package via the following:

    go get -u github.com/rafos/go-multimap

## Usage ##
The go-multimap package can be used similarly to the following:
```go
package main

import (
	"fmt"
	"github.com/rafos/go-multimap/slicemultimap"
)

func main() {
	usPresidents := []struct {
		firstName  string
		middleName string
		lastName   string
		termStart  int
		termEnd    int
	}{
		{"George", "", "Washington", 1789, 1797},
		{"John", "", "Adams", 1797, 1801},
		{"Thomas", "", "Jefferson", 1801, 1809},
		{"James", "", "Madison", 1809, 1817},
		{"James", "", "Monroe", 1817, 1825},
		{"John", "Quincy", "Adams", 1825, 1829},
		{"John", "", "Tyler", 1841, 1845},
		{"James", "", "Polk", 1845, 1849},
		{"Grover", "", "Cleveland", 1885, 1889},
		{"Benjamin", "", "Harrison", 1889, 1893},
		{"Grover", "", "Cleveland", 1893, 1897},
		{"George", "Herbert Walker", "Bush", 1989, 1993},
		{"George", "Walker", "Bush", 2001, 2009},
		{"Barack", "Hussein", "Obama", 2009, 2017},
	}

	m := slicemultimap.New[string, string]()

	for _, president := range usPresidents {
		m.Put(president.firstName, president.lastName)
	}

	for _, firstName := range m.KeySet() {
		lastNames, _ := m.Get(firstName)
		fmt.Printf("%v: %v\n", firstName, lastNames)
	}
}
```

Example output:
```sh
$ go run example/example.go
John: [Adams Adams Tyler]
Thomas: [Jefferson]
James: [Madison Monroe Polk]
Grover: [Cleveland Cleveland]
Benjamin: [Harrison]
Barack: [Obama]
George: [Washington Bush Bush]
```

## Benchmarks ##
To see the benchmark, run the following on each of the sub-packages:

`go test -run=NO_TEST -bench . -benchmem  -benchtime 1s ./...`
<pre>
goos: darwin
goarch: amd64
pkg: github.com/rafos/go-multimap/slicemultimap
cpu: Intel(R) Core(TM) i7-8559U CPU @ 2.70GHz
BenchmarkMultiMapGet100-8                 578694              2001 ns/op               0 B/op          0 allocs/op
BenchmarkMultiMapGet1000-8                 48384             24430 ns/op               0 B/op          0 allocs/op
BenchmarkMultiMapGet10000-8                 3927            315908 ns/op               0 B/op          0 allocs/op
BenchmarkMultiMapGet100000-8                 298           3822512 ns/op               0 B/op          0 allocs/op
BenchmarkMultiMapPut100-8                 209097              6214 ns/op            8590 B/op          0 allocs/op
BenchmarkMultiMapPut1000-8                 20220             64400 ns/op           84408 B/op        745 allocs/op
BenchmarkMultiMapPut10000-8                 1525            715396 ns/op          596457 B/op       9822 allocs/op
BenchmarkMultiMapPut100000-8                 128          10531982 ns/op         7172956 B/op     105994 allocs/op
BenchmarkMultiMapPutAll100-8              224827              5502 ns/op            7989 B/op          0 allocs/op
BenchmarkMultiMapPutAll1000-8              18633             67812 ns/op           91091 B/op        745 allocs/op
BenchmarkMultiMapPutAll10000-8              1554            726543 ns/op          586781 B/op       9821 allocs/op
BenchmarkMultiMapPutAll100000-8              128          10541793 ns/op         7172955 B/op     105994 allocs/op
BenchmarkMultiMapRemove100-8              594860              1976 ns/op               0 B/op          0 allocs/op
BenchmarkMultiMapRemove1000-8              41503             30151 ns/op            5952 B/op        744 allocs/op
BenchmarkMultiMapRemove10000-8              2859            353375 ns/op           77952 B/op       9744 allocs/op
BenchmarkMultiMapRemove100000-8              334           3206979 ns/op          797956 B/op      99744 allocs/op
BenchmarkMultiMapRemoveAll100-8          1905345               630.3 ns/op             0 B/op          0 allocs/op
BenchmarkMultiMapRemoveAll1000-8          187369              6343 ns/op               0 B/op          0 allocs/op
BenchmarkMultiMapRemoveAll10000-8          18818             62743 ns/op               0 B/op          0 allocs/op
BenchmarkMultiMapRemoveAll100000-8          1756            635113 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/rafos/go-multimap/slicemultimap      30.781s
</pre>
