package main

import "fmt"

type RowDb struct {
	amount float64
	atype  string
}

func main() {
	//var data map[string]RowDb
	/*
		data := make(map[string]RowDb)
		data["c"] = RowDb{0.7, "3"}
		row := data["c"]
		row.amount += 3.4
		fmt.Print(data, row)
		var a interface{}
		a = &A{}
		fmt.Println(a == nil)
		b := &B{}
		fmt.Println(b == nil)
	*/

	isli := []int{1, 2}
	//bsli := []int{}
	var csli = []int{}
	fmt.Println(csli == nil, csli)
	fmt.Println(append(isli, csli...))

}

type I interface{}
type A struct {
}
type B struct {
}
