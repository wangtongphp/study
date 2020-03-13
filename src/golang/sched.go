package main

import (
	"fmt"
	"runtime"

	//	"runtime"
	"sync/atomic"
	"time"
)

var count int64

func main() {

	fmt.Println("vim-go")

	for {
		go test()
		time.Sleep(time.Second)
		runtime.GC()
	}



	for {
		time.Sleep(time.Second)
	}

}

func test() {

	atomic.AddInt64(&count, 1)
	defer atomic.AddInt64(&count, -1)

	//runtime.LockOSThread()
	//defer runtime.UnlockOSThread()

	time.Sleep(time.Second * 10)

}
