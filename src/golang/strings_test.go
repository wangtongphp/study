/*
wangtong1v:~/devspace/study/src/golang $  go test -race -v strings_test.go -run ""  -bench="B"
=== RUN   TestSearch
--- PASS: TestSearch (0.00s)
strings_test.go:29: true
strings_test.go:30: 0
BenchmarkEEqualhmarkSearchIndex-8      50000000        39.1 ns/op
BenchmarkSearchContains-8   20000000        66.6 ns/op
BenchmarkEEqual-8           100000000        14.6 ns/op
BenchmarkENottrue-8         100000000                 14.7 ns/op
PASS
ok  command-line-arguments7.367s
*/

// strings.Index 比strings.Contains性能高一倍。
package main

import (
	"strings"
	"testing"
)

// 3.53 ns/op
func BenchmarkSearchIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if strings.Index("hello world", "r") > -1 {

		}
	}
}

// 6.60 ns/op
func BenchmarkSearchContains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if strings.Contains("hello world", "r") {

		}
	}
}

func TestSearch(t *testing.T) {

	t.Log(strings.Contains(" hello world", ""))
	t.Log(strings.Index(" hello world", ""))
}

func BenchmarkEEqual(b *testing.B) {
	var abc bool = false
	for i := 0; i < b.N; i++ {
		if abc == false {

		}
	}
}

func BenchmarkENottrue(b *testing.B) {
	var abc bool = false
	for i := 0; i < b.N; i++ {
		if !abc {

		}
	}
}
