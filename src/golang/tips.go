//https://chai2010.gitbooks.io/advanced-go-programming-book/content/appendix/appendix-a-trap.html
package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"
)

func main() {
	// main1()
	// fmt.Println("main2", main2())
	// main3()
	// main4()
	// main5()
	// main6()
	main7()
}

//数组是值传递, 切片是引用传递
func main1() {
	x := []int{1, 2, 3}
	func(arr []int) {
		arr[0] = 7
		fmt.Println(arr)
	}(x)
	fmt.Println(x)
	//out:
	//[7 2 3]
	//[7 2 3]

	//*******************
	// 这里明确个数后就是值传递了
	y := [3]int{1, 2, 3}
	func(arr [3]int) {
		arr[0] = 7
		fmt.Println(arr)
	}(y)
	fmt.Println(y)
	//out:
	// [7 2 3]
	// [1 2 3]
}

func main2() (err error) {
	// 若这里是:=, 则报错err is shadowed during return
	if _, err = strconv.Atoi("a"); err != nil {
		return
	}
	return
}

func main3() {
	//错误。 半数概率，输出前就终止了
	go println("hello")
	runtime.Gosched()
}

func main4() {
	runtime.GOMAXPROCS(1)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()

	for {
		//Gosched若没有， 则会出现， 独占CPU导致其它Goroutine饿死
		runtime.Gosched()
	} // 占用CPU
}

//不同Goroutine之间不满足顺序一致性内存模型
func main5() {
	runtime.GOMAXPROCS(1)
	// var msg string
	// var done bool = false
	msg := ""
	done := false

	// TODO 为何不用传参进去呢?
	go func() {
		time.Sleep(time.Second * 1)
		msg = "hello, world"
		done = true
	}()

	for {
		if done {
			println(msg)
			break
		}
		// 必须有Sleep,
		time.Sleep(time.Second * 0)
	}
}

//在循环内部执行defer语句, 解决的方法可以在for中构造一个局部函数，在局部函数内部执行defer：
func main6() {
	for i := 0; i < 5; i++ {
		f, err := os.Open("/path/to/file")
		if err != nil {
			log.Fatal(err)
		}
		// 错误。 defer在函数退出时才能执行，在for执行defer会导致资源延迟释放：
		defer f.Close()
	}
}

func main7() {
	return
}
