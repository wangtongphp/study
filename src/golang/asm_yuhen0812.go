package main

import (
	"fmt"
	"time"
	"unsafe"
)

func main() {
	test6()
}

// go build -gcflags "-N -l" asm_yuhen0812.go ; go tool objdump -s "main\.test1" asm_yuhen0812
//go:noinline
//go:nosplit
func test1() {
	var i int = 33 // 0x21
	var p *int = &i
	println(i, p, *p)
}

//go:noinline
//go:nosplit
func test2() {
	var i int = 0x01020304056677ff
	var p *[8]byte
	p = (*[8]byte)(unsafe.Pointer(&i))

	//p[0] = 5
	//p[5] = 6

	fmt.Printf("%v, %016X, \n", p, i)
}

//go:noinline
//go:nosplit
func test3() {

	x := [8]byte{1, 2, 3, 4}
	var p *byte = &x[0]

	u := uintptr(unsafe.Pointer(p))
	fmt.Printf("%p, %x, %d \n", p, u, *p)

	u++
	fmt.Printf("%p, %x, %x \n", p, u, *p)

	p2 := (*byte)(unsafe.Pointer(u))
	*p2 += 100
	fmt.Println(p2, *p2, x)

}

//go:noinline
//go:nosplit
func test4() {
	d := []byte{'a', 'b', 'c', 'd'}
	s := string(d) // copy
	fmt.Println(d, s)

	type str struct {
		data uintptr
		len  int
	}

	type sli struct {
		data uintptr
		len  int
		cap  int
	}

	dh1 := *(*sli)(unsafe.Pointer(&d))
	dh2 := *(*str)(unsafe.Pointer(&d))
	dh3 := *(*string)(unsafe.Pointer(&d))

	sh1 := *(*sli)(unsafe.Pointer(&s))
	sh2 := *(*[]byte)(unsafe.Pointer(&s))

	fmt.Printf("%#v, %#v, %#v, %#v, %#v \n", dh1, dh2, dh3, sh1, sh2)
}

// overflow
//go:noinline
//go:nosplit
func test5() {
	a, b := 1, 2
	fmt.Println(a, b)

	p := &a
	u := uintptr(unsafe.Pointer(p)) + unsafe.Sizeof(b)

	p = (*int)(unsafe.Pointer(u))
	*p += 100
	fmt.Println(a, b, p, *p)
}

// 不同指针类型对GC影响
// *data, unsafe.Pointer, uintptr, &data.x
//➜  golang git:(master) ✗ go build asm_yuhen0812.go && GODEBUG=gctrace=1 ./asm_yuhen0812
//gc 1 @0.000s 2%: 0.005+0.14+0.002 ms clock, 0.020+0.042/0.025/0.12+0.008 ms cpu, 10->10->10 MB, 11 MB goal, 4 P
//gc 2 @1.002s 0%: 0.004+0.12+0.003 ms clock, 0.016+0/0.081/0.13+0.014 ms cpu, 20->20->20 MB, 21 MB goal, 4 P
type data struct {
	x [10 << 20]byte
	y int
}

func test6() {
	d := make([]*data, 100)
	for {
		d = append(d, test6_1())
		time.Sleep(time.Second * 1)
	}
	fmt.Println("done")
}

func test6_1() *data {

	d := new(data)
	d.x[0] = 100

	return d
	//return unsafe.Pointer(d)
	//return uintptr(unsafe.Pointer(d))
}

//go:noinliine
//go:nosplit
func xhh() {

	a := 1
	v := []*int{&a}

	vv := map[string]*int{"a": &a}

	fmt.Println(v[0], vv["a"])

	fmt.Printf("%v, %#v", v[0], unsafe.Pointer(vv["a"]))

}
