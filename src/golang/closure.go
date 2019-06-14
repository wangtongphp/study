package main

import "fmt"

func main() {
	f := fibonacci()
	for i := 0; i < 11; i++ {
		fmt.Println(f())
	}
}

func fibonacci() func() int {
	var a0, a1 = 0, 1
	return func() int {
		a1, a0 = a1+a0, a1
		return a0
	}
}
