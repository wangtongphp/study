package main

import (
    "fmt"
    )

func main(){
     //var vala Intcus = 3
     //vala.Add(3)
     //fmt.Println(vala, vala.Less(34))
     //quoteT()
     o1 := new(C1)
     o2 := &C1{"wangtong"}
     o3 := &C2{Age: 22}
     o1.Add()
     o2.Add()
     o3.Add()
}

type C1 struct{
    Name string
}

func(a *C1) Add(){
    fmt.Println("this is C1.Add",a.Name)
}

func(a *C1) Del(){
    fmt.Println("this is C1.Del")
}

type C2 struct{
    Age int
    *C1
}

func(a *C2) Add(){
    fmt.Println("this is C2.Add",a.Name, a.Age, a.Del())
}


//Intcus 对象
type Intcus int

//Intcus对象的Add方法
func(a *Intcus) Add(b Intcus) {
    *a += b
}

//Intcus对象的Less方法
func(c Intcus) Less(b Intcus) bool {
    return (c < b)
}

//slice 为引用类型
func quoteT(){
    arr_v:= [4]int{3,6,3,2}
    slic_v := arr_v[:]
    slic_v[1] = 9
    fmt.Println(arr_v) //[3 9 3 2]
}

