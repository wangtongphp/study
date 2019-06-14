package main

import (
	"fmt"
	//"reflect"
)

func main() {
	datarange()
	slirange()
	fmt.Println("\n")
	sliappend()
}
func datarange() {

	m := map[int]int{
		1: 11,
		2: 22,
	}
	// 遍历只接一个值，则是key，非value
	for i := range m {
		fmt.Println(i)
	}
}

func slirange() {

	s := []int{1, 2, 3, 4, 5}
	fmt.Println(len(s), cap(s), fmt.Sprintf("%p", s[0]))
	for i, v := range s {
		s[3] = 333
		fmt.Println(len(s), cap(s), fmt.Sprintf("%p", s[0]))
		s = append(s, v+5)
		fmt.Println(cap(s), i, v)
	}
	fmt.Println(s)
}

//reslice 新建切片依旧指向原底层数组， 修改对所有关联切片可见 append 超出sli的cap则会新分配空间，不会改变原数组
func sliappend() {
	arr := []int{1, 2, 3, 4, 5}
	s := arr[:]
	s1 := arr[:]
	//三个[0]指针地址相等
	fmt.Println(len(s), cap(s), fmt.Sprintf("%p %p %p", &arr[0], &s[0], &s1[0]))

	//s地址变化
	s = append(s, 6)
	fmt.Println(len(s), cap(s), fmt.Sprintf("%p %p %p", &arr[0], &s[0], &s1[0]))

	//s1会同时改变arr的值, s因为是新空间，所以不会影响其他slice的值
	s1[1] = 222
	s[1] = 22
	fmt.Println(arr, s, s1, len(s), cap(s), fmt.Sprintf("%p %p", &s, &s[0]))

	//s的cap又变了，s[0]地址变了
	s = append(s, 7, 8, 9, 10, 11)
	fmt.Println(len(s), cap(s), fmt.Sprintf("%p %p", &s, &s[0]))
}
