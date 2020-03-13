package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 4)
	go consumer(ch)
	go producter(ch)
	time.Sleep(3 * time.Second)
}

func consumer(ch <-chan int) {
	for k := range ch {
		fmt.Println(k)
	}
	fmt.Println("consumer done")
}

func producter(ch chan<- int) {
	defer func() {
		close(ch)
	}()
	for i := 0; i < 311; i++ {
		ch <- i
	}
}
