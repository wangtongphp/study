package main

import "fmt"

// map race 无法recover
func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover", err)
		}
	}()
	m := map[int]int{
		1: 1,
		2: 2,
	}
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("recover", err)
			}
		}()
		for i := 0; i < 1111; i++ {
			v, ok := m[i]
			if ok != true || v == 0 {
				fmt.Println("err")
			}
		}
	}()
	for i := 0; i < 1111; i++ {
		m[i] = i
	}

	fmt.Println("vim-go",
		string(byte(01)), byte(01))

}
