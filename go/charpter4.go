package main

import (
    "fmt"
    "runtime"
    )

func main(){
    fmt.Println(runtime.NumCPU())
    runtime.GOMAXPROCS(runtime.NumCPU())
    var c chan int 
    for i:=1; i<9; i++{
        c <- i
    }
    res := <- c
    fmt.Println(res)
}


