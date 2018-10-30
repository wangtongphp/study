package main

import (
    "io"
    "fmt"
    "os"

    )

func main(){
    ioT()
}

func ioT(){
    
    w,err := io.Copy("/tmp/t1.txt","/tmp/t2.txt")
    fmt.Println(w,err)
    
}
