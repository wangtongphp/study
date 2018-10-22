package main

import (
	"fmt"
)

func main() {
	//first-class object
	//var f1 = f
	//f1()

	//escape analysis
	var a *int = escapeAnalysis()
	fmt.Println(a, *a)
}

func f() {
	fmt.Println("f()")
}

func escapeAnalysis() *int {
	var a = 999
	return &a
}
