//并发
package main

import (
	"testing"
	// "sync"
)

// go test -v go-yuhen-8-goroutine_test.go -run=NONE  -bench="Perf1" -benchtime=5s -benchmem -gcflags "-N -l"
func BenchmarkPerf1(b *testing.B) {
	var (
		max     = b.N
		bufsize = 100
	)
	done := make(chan struct{})
	data := make(chan int, bufsize)

	//consumer
	go func() {
		cnt := 0
		for i := range data {
			cnt += i
		}
		close(done)
		b.Log(b.N, cnt)
	}()

	//producer
	for i := 0; i < max; i++ {
		data <- i
	}
	close(data)

	<-done
	b.Log(b.N, "done")
}

//test1
func BenchmarkPerf2(b *testing.B) {
	const (
		block   = 500
		bufsize = 100
	)
	var max = b.N

	done := make(chan struct{})
	data := make(chan [block]int, bufsize)

	//consumer
	go func() {
		cnt := 0
		for i := range data {
			for j := range i {
				cnt += j
			}
		}
		close(done)
		b.Log(b.N, cnt)
	}()

	//producer
	for i := 0; i < max; i += block {
		var b [block]int
		for j := 0; j < block; j++ {
			b[j] = i + j
			if i+j == max-1 {
				break
			}
		}
		data <- b
	}
	close(data)

	<-done
	b.Log(b.N, "done")
}
