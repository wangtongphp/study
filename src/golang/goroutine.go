package main

import (
	"context"
	"errors"
	"fmt"
	// "fmt"
	"sync"
	"sync/atomic"
	"time"

	// "testing"
	// "time"
	"golang.org/x/sync/errgroup"
)

func main() {
	gf1()

}

func f0(){

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

func gf1(){
	//var g errgroup.Group
	g,ctx:=errgroup.WithContext(context.Background())

	//res := make([]int,0)
	res := make([]int,0,20000)
	res1:=0
	var res2 int32
	g.Go(func() error{
		time.Sleep(10*time.Nanosecond)
		for i:=0;i<10000;i++ {
			res = append(res,i)
			res1++
			atomic.AddInt32(&res2, +1)
		}
		return nil
		return errors.New("g1")
	})
	g.Go(func() error{
		for i:=10000;i<20000;i++ {
			res = append(res,i)
			res1++
			atomic.AddInt32(&res2, +1)
		}
		//time.Sleep(1*time.Second)
		//return nil
		return errors.New("g2")
	})

	 <-ctx.Done()
	 //err:=g.Wait()
	 //fmt.Println(err)


	fmt.Println(len(res))
	fmt.Println(res1)
	fmt.Println(res2)

}
