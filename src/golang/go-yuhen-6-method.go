package main

import (
	"fmt"
)

type N int

func (n N) test() {
	fmt.Println(n)
}

func main() {
	var n N = 25

	f1 := N.test
	f1(n)
}
