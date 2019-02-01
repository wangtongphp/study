/*
wangtong1v:~/devspace/study/src/golang $  go test go-yuhen-5-data_test.go -bench . -benchmem
BenchmarkArray-8     1000000          1286 ns/op           0 B/op          0 allocs/op
BenchmarkSlice-8     1000000          2104 ns/op        8192 B/op          1 allocs/op
PASS
ok      command-line-arguments  3.437s
*/
package main

import (
	"testing"
)

func array() [1024]int {
	var x [1024]int
	for i := 0; i < len(x); i++ {
		x[i] = i
	}
	return x
}

func slice() []int {
	x := make([]int, 1024)
	for i := 0; i < len(x); i++ {
		x[i] = i
	}
	return x
}

func BenchmarkArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		array()
	}
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice()
	}
}
