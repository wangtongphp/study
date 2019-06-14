package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover....", err)
		}
	}()
	fmt.Println("vim-go")
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("recover....", err)
			}
		}()
		panic("aaa")
	}()

	time.Sleep(time.Second * 1)

}
