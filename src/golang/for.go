package main

import "fmt"

func main() {
	//break 仅一层,多层用GOTO
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Println(i, j)
			if j == 1 {
				break
			}
		}
	}
}
