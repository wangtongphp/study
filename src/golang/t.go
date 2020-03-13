package main

import (
	"fmt"
)

func main() {
	a, b := 16, 17
	a, b = b, a
	fmt.Println(a, b)
	/*
		m := map[int]int{
			0:  0,
			1:  1,
			2:  2,
			3:  3,
			4:  4,
			5:  5,
			6:  6,
			7:  7,
			8:  8,
			9:  9,
			17: 33,
		}

		p := unsafe.Pointer(&m)
		fmt.Println(m[17], p)

		fmt.Println("0.0.2" >= "0.0.2")
	*/
}

/*
1
10
11
100
101
111
1000
*/
