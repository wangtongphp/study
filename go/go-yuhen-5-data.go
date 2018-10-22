package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {

	var bs []byte = []byte(" byte ")
	fmt.Println(toString(bs), string(bs))

	isRune()

	isSli()

	sliPtr()

	structDef()
}

func toString(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

//单引号双引号默认类型不同
func isRune() {
	r := '我'
	fmt.Printf("%T, %v\n", r, r)

	s := "我"
	fmt.Printf("%T, %v\n", s, s)
}

//默认类型是数组还是切片
func isSli() {
	a := [...]int{0, 1, 2}
	b := []int{0, 1, 2}
	fmt.Printf("a:%T, %v\n", a, a)
	fmt.Printf("b:%T, %v\n", b, b)
}

//slice 指针
func sliPtr() {
	a := []int{0, 1, 2}
	p := &a
	(*p)[0] += 100
	fmt.Printf("a, p ,%v, %v\n", a, p)
}

//struct define
func structDef() {
	type u struct {
		name string
		addr struct {
			prov string
			area string
		}
		string
		int
	}

	a := u{
		name: "wang",
		// addr: {
		// 	prov: "shandong",
		// 	area: "",
		// },
		string: strconv.Itoa(3),
		int:    3,
	}

	a.addr.prov = "shandong"

	fmt.Printf("%#v, %v", strconv.Atoi("a"))
}
