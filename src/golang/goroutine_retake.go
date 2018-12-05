// wangtong1v:~/devspace/study/src/golang $  GODEBUG="schedtrace=1000,scheddetail=1" ./goroutine_retake

package main

import (
	"log"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	go printNum()
	ret := 0
	for i := 0; ; i++ {
		ret += add(i)
	}
	log.Println(ret)
}

func printNum() {
	for i := 0; i < 9999999; i++ {
		log.Println(i * 1)
	}
}

func add(i int) int {
	i += 10
	return i
}
