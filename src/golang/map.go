package main

import (
	"fmt"
)

func main() {
	//data:=[6]int{1,5,2,7,2,4}
	//exec1(data, 6, 2)
	var s = SSMap{"d": "a"}
	s.f1()
}

type SSMap map[string]string

func (m SSMap) f1() {

	m.f2()
	m["a"] = "1"
	m["b"] = "2"

	fmt.Println(m, len(m))
}

func (m SSMap) f2() {
	m["c"] = "c"

}
