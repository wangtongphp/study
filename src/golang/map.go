package main

import (
	"fmt"
)

type SSMap map[string]string

func main() {

	//data:=[6]int{1,5,2,7,2,4}
	//exec1(data, 6, 2)

	//var s = SSMap{"d": "a"}
	//s.f1()

	mapokidiom()

	concur()

	// 多维
	//fmore()
}

func concur() {
	m := make(map[int]int, 8)

	go func() {
		m[1] = 1
		fmt.Println(m[1])
	}()
	fmt.Println(m[1])

	select {}
}

func fmore() {
	var QuestionData map[int]map[int]*struct{}
	QuestionData[1][2] = &struct{}{}
	fmt.Println(QuestionData)
}

// map引用类型, 会改变参数的值
func (m SSMap) f1() {

	m.f2()
	m["a"] = "1"
	m["b"] = "2"

	fmt.Println(m, len(m))
}

func (m SSMap) f2() {
	m["c"] = "c"
}

// map.go:36:15: invalid operation: b[2] (type int does not support indexing)
func mapokidiom() {
	a := map[int]int{1: 11}
	fmt.Println(a[1])
	//fmt.Println(b[2])
}

/*
func ChangeFoo(f *runtime.hmap) {
	p := unsafe.Pointer(f)
	// 事先获取或者通过 reflect获得
	// 本例中是第一个字段，所以offset=0
	offset := uintptr(0)
	ptr2x := (*int)(unsafe.Pointer(uintptr(p) + offset))
	fmt.Println(*ptr2x)
	*ptr2x = 100
	fmt.Println(f.X())
}
*/
