package main

import (
	//"regexp"
	"fmt"
	"strings"
)

func main() {
	grepT()
	split()
}

func grepT() {
	str := []rune("abcde")
	for _,v:=range str{
		fmt.Println(v)
	}
	fmt.Println(str[3:4])
}

func split() {
	s := strings.Split("9a", "")
	fmt.Println(s)
}
