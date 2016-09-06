package main

import (
    "fmt"
    //"runtime"
    //"strconv"
)

type ByteSize float64

const (
    _ = iota
    KB ByteSize = 1 << (iota * 10)
    MB
    GB
    TB
)


type Xm struct {
    City string
    Name string
}

func testFm(format string, args ...interface{}){
    fmt.Println(args)
    s := fmt.Sprintf(format, args)
    fmt.Println(s)
}

func closeure(x int) func(int) int {
    fmt.Println(&x)
    return func(y int) int {
        fmt.Println(&x)
        return x+y
    }
}

func res1()(res int){
    defer func(){
        res = res + 5
    }()
    return 1
}

func func1(c chan bool, index int){
    a := 1
    for i:=0; i<999999999; i++ {
        a += 1
    }
    fmt.Println(index, a)
    c <- true 
}

func main() {
    /*
    runtime.GOMAXPROCS(runtime.NumCPU())
    c := make(chan bool)
    for i:=0; i<10; i++ {
        go func1(c, i)
    }
    for i:=0; i<10; i++ {
        <- c
    }

    f := res1()
    fmt.Println(f)
    
    //panic("m_b")
    defer fmt.Println("o1")
    defer fmt.Println("o2")
    defer fmt.Println("o3")
    //panic("m_e")

    f := closeure(33)
    fmt.Println(f(22))
    fmt.Println(f(222))

    ft := testFm
    ft("HH",34,55)
    testFm("%s hehe %f", "haha", 22.2)
    var testSlice = []interface{}{33,22}
    fmt.Println(testSlice)
    fmt.Println(testSlice...)
    
    var a_stru = &Xm{
        City: "BJ",
        Name: "wt",
    }
    a_stru.City = "SH"
    fmt.Println(a_stru)

    var i_map = map[string]string{"k_a":"v_a","k_b":"v_b"}
    for k, v := range i_map{
        fmt.Println(k,v)
    }

    var mi_map = make(map[string]string)
    if _, ok := mi_map["k"]; !ok{
        mi_map["k"] = "3"
    }
    fmt.Println(mi_map)

    xiaomi := new(Xm)
    i := new(int)
    sl := new([]string)
    fmt.Println(xiaomi, i, sl)

    x_slice := make([]int, 10)
    x_map := make(map[string]int, 10)
    x_channel := make(chan string, 10)
    fmt.Println(x_slice, x_map, x_channel)

    fmt.Println("hello world!\n",runtime.Version())
    */
}

