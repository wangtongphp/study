package main

import (
	"fmt"
	"strconv"
)

func main(){
f2()
}

func f1(){
	a:=1.5
	r:=strconv.FormatFloat(a*0.7, 'f', -1, 64)
	fmt.Println(r)

	n:=fmt.Sprintf("%.2f\n",a*0.7)
	m:=strconv.FormatFloat(a*0.7, 'f', 2, 64)
	fmt.Println(n,m)
}

func f2(){
	a:="0.0"
	//format(x,'.2%')
	b,e:=strconv.Atoi(a)
	c,e2:= strconv.ParseFloat(a,64)
	fmt.Println(b,e)
	fmt.Println(c<=0.0,e2)
}