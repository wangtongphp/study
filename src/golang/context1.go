package main

import (
	"context"
	"fmt"
	"time"
)

// 很难适用于ontp场景， 除非有for下面执行多个操作， 或者是以 <-ctx.Done() 作为埋点条件各处进行判断
func work(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			println(name, " get message to quit")
			return
		default:
			//println(name, "1 is running", time.Now().String())
			//time.Sleep(time.Second * 5)
			//println(name, "2 is running", time.Now().String())

		}
	}
}

func main() {

	/*
		//cancel
		ctx, cancel := context.WithCancel(context.Background())
		go work(ctx, "work1")

		time.Sleep(time.Second * 3)
		cancel()
		time.Sleep(time.Second * 5)
		   // with value
		   ctx1, valueCancel := context.WithCancel(context.Background())
		   valueCtx := context.WithValue(ctx1, "key", "test value context")
		   go workWithValue(valueCtx, "value work", "key")
		   time.Sleep(time.Second * 3)
		   valueCancel()
	*/
	// timeout
	//timeCancel()及timeout都会触发<-ctx.Done为true
	ctx2, timeCancel := context.WithTimeout(context.Background(), time.Microsecond*1)
	defer timeCancel()
	fmt.Println("0")
	go work(ctx2, "time cancel")
	fmt.Println("1")
	time.Sleep(time.Second * 5)
	fmt.Println("2")
	/*
	   // deadline
	   ctx3, deadlineCancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*3))
	   go work(ctx3, "deadline cancel")
	   time.Sleep(time.Second * 5)
	   deadlineCancel()

	   time.Sleep(time.Second * 3)
	*/

}

func workWithValue(ctx context.Context, name string, key string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Value(key))
			println(name, " get message to quit")
			return
		default:
			println(name, " is running", time.Now().String())
			time.Sleep(time.Second)
		}
	}
}
