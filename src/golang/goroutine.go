package main

import (
	// "fmt"
	"sync"
	// "testing"
	// "time"
)

func main() {

	//producer
	// c := make(chan int, 256)
	var wg sync.WaitGroup
	// go func() {
	// 	for i := 0; i < 123456789; i++ {
	// 		c <- i
	// 	}
	// 	close(c)
	// }()

	//time.Sleep(time.Second * 1)
	//fmt.Println("c:", len(c), cap(c), b.N)

	//consumer
	// for i := 0; i < 1; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		for {
	// 			i, ok := <-c
	// 			if ok == false {
	// 				break
	// 			}
	// 			_ = i + i*2
	// 		}
	// 		wg.Done()
	// 	}()
	// }
	go func() {
		wg.Add(1)
		for i := 3; ; i++ {
			i++
		}
		wg.Done()
	}()

	for i := 3; ; i++ {
		i++
	}
	wg.Wait()

}
