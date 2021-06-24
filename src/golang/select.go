package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	//d := make(chan int)
	select {
	case <-c:
		fmt.Println("c")
	//case <-d:
	//	fmt.Println("d")
	case <- time.After(time.Second*4):
		fmt.Println("timeout")
	}

	go func() {
		time.Sleep(time.Second*3)
		c <- 1
	}()

	//d <- 2
	//time.Sleep(time.Second*3)
	fmt.Println("done")
}
