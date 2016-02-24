package main

import (
    "fmt"
    )

func main(){
    var c chan int 
    for i:=1; i<9; i++{
        c <- i
    }
    res := <- c
    fmt.Println(res)
}
