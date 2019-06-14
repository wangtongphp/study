package main

import "fmt"

func main() {
	var tt interface{}
	tt = nil
	t1, ok := tt.(string)
	//t1 := tt.(string)
	fmt.Println(t1, ok)
}
