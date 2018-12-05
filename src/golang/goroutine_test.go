// go test -race -v goroutine_test.go -run "Cl"  -bench="xxx"
package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_Closure(t *testing.T) {
	s := [...]string{
		"first",
		"second",
		"third",
	}

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(s))

	for _, item := range s {
		go func(i string) {
			// t.Log(i)
			t.Log(item)
			waitGroup.Done()
		}(item)
		time.Sleep(1 * time.Microsecond)
	}

	waitGroup.Wait()
}

/**
* @brief go routine, chan
*
* @param testing.T
*
* @return
 */
func Test_Chan(t *testing.T) {

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	c := make(chan int, 3)

	//producer
	go func() {
		for {
			for _, v := range data {
				c <- v
			}
			time.Sleep(time.Second * 10)
		}
	}()

	time.Sleep(time.Second * 1)
	fmt.Println("c:", len(c), cap(c))

	//consumer
	for i := 0; i < 2; i++ {
		go func() {
			for {
				time.Sleep(time.Second * 1)
				select {
				case e := <-c:
					fmt.Println("item:", e)
				case <-time.After(2 * time.Second):
					fmt.Println("timeout")
				}
			}
		}()
	}

	time.Sleep(time.Second * 1)
	fmt.Println(" end ")
	select {}
}

//go routine, sync.WaitGroup
func Test_Wg(t *testing.T) {

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	c := make(chan int, 3) //并发数
	var wg sync.WaitGroup

	time.Sleep(time.Second * 1)
	fmt.Println("c:", len(c), cap(c))

	for _, v := range data {
		wg.Add(1)
		c <- 1
		go func(v int) {
			defer func() {
				wg.Done()
				<-c
			}()
			time.Sleep(time.Second * 2)
			fmt.Println(" item: ", v)
		}(v)
	}
	wg.Wait()

	//把chan填满也能实现wg.Wait()同样效果, 但变量未释放
	// for i := 0; i < 3; i++ {
	// 	c <- 1
	// }

	//time.Sleep(time.Second * 1)
	fmt.Printf(" end , %v", wg)
}

/**
* @brief channel
Chan1 始终都是10个go routine在跑
Chan2 起b.N个go routine跑

dev10:~/devspace/study/go $  go test -v goroutine_test.go -run "WF"  -bench="Chan" -benchtime=5s
goos: linux
goarch: amd64
Benchmark_Chan1-8    5000000      1039 ns/op
Benchmark_Chan2-8    5000000      1048 ns/op
PASS
ok  command-line-arguments30.638s

* @return
*/
func Benchmark_Chan1(b *testing.B) {

	//producer
	c := make(chan int, 256)
	var wg sync.WaitGroup
	go func() {
		for i := 0; i < b.N; i++ {
			c <- i
		}
		close(c)
	}()

	//time.Sleep(time.Second * 1)
	//fmt.Println("c:", len(c), cap(c), b.N)

	//consumer
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			for {
				i, ok := <-c
				if ok == false {
					break
				}
				_ = i + i*2
			}
			wg.Done()
		}()
	}

	wg.Wait()

}

/**
* @brief
*
* @param testing.B
*
* @return
 */
func Benchmark_Chan2(b *testing.B) {

	//producer
	c := make(chan int, 10)
	var wg sync.WaitGroup
	go func() {
		for i := 0; i < b.N; i++ {
			c <- i
		}
	}()

	//time.Sleep(time.Second * 1)
	//fmt.Println("c:", len(c), cap(c), b.N)

	ch := make(chan int, 10)
	//consumer
	for i := 0; i < b.N+1; i++ {
		wg.Add(1)
		ch <- 1
		go func() {
			defer func() {
				wg.Done()
				<-ch
			}()
			_, ok := <-c
			if ok == false {
				b.Log("ok == false")
			}
		}()
	}
	wg.Wait()

}

/**
* @brief
*
* @param testing.B
*
* @return
 */
func Benchmark_Wg(b *testing.B) {

	//producer
	data := make([]int, 0)
	for i := 0; i < b.N; i++ {
		data = append(data, i)
	}

	c := make(chan int, 10)
	var wg sync.WaitGroup

	for _, v := range data {
		wg.Add(1)
		c <- 1
		go func(v int) {
			defer func() {
				wg.Done()
				<-c
			}()
			// ...
		}(v)
	}

	//把chan填满也能实现wg.Wait()同样效果
	// for i := 0; i < 3; i++ {
	// 	c <- 1
	// }
	wg.Wait()

	//fmt.Printf(" end , %v", c)
}
