package main

import "fmt"

func main() {
	c := make(chan int)
	d := make(chan int)
	go func() {
		select {
		case <-c:
			fmt.Println("c")
		case <-d:
			fmt.Println("d")

		}
	}()
	c <- 1
	d <- 2
	fmt.Println("done")
}
