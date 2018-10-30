//并发
package main

import (
	"fmt"
	"testing"
	// "sync"
)

//test1
func BenchmarkPerf1(b *testing.B) {
	var (
		max = b.N
		//block   = 500
		bufsize = 100
	)
	done := make(chan struct{})
	data := make(chan int, bufsize)
	go func() {
		defer close(done)
		cnt := 0
		for i := range data {
			cnt += i
		}
	}()

	for i := 0; i < max; i++ {
		data <- i
	}
	close(data)
	<-done
	fmt.Println("done")
}

//test1
func BenchmarkPerf2(b *testing.B) {
	var (
		max     = b.N
		block   = 500
		bufsize = 100
	)
	done := make(chan struct{})
	data := make(chan [block]int, bufsize)
	go func() {
		defer close(done)
		cnt := 0
		for i := range data {
			cnt += i
		}
	}()

	for i := 0; i < max; i++ {
		data <- i
	}
	close(data)
	<-done
	fmt.Println("done")
}
