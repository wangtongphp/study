// strings.Index 比strings.Contains性能高一倍。
package main

import (
	"strings"
	"testing"
)

// 3.53 ns/op
func BenchmarkSearchIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if strings.Index("hello world", "") > -1 {

		}
	}
}

// 6.60 ns/op
func BenchmarkSearchContains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if strings.Contains("hello world", "") {

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
