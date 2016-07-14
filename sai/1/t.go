package main

import (
	"fmt"
)

func main() {
	sli := make([]int, 50)
	sli[40] = 3333
	fmt.Print(sli[33])
}
