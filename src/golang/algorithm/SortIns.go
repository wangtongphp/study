package main

import (
	"fmt"
	"sort"
)

func main() {
	s:=[]int{1,3,4,2,1,5,9,2,6,111,33,44,22}
	//sort.Ints(s)
	//sort.Ints(s)
	sort.Slice(s, func(i, j int) bool {
		return s[i]>s[j]
	})
	fmt.Println(s)
}
