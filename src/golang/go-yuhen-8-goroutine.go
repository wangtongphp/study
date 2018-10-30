//并发
package main

import (
	"fmt"
	// "sync"
)

func main() {
	chan1()
}

//test1
func chan1() {
	var (
		max = 5000000
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
