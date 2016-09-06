package main

import (
    "fmt"
    )

func main(){
    SwitchT()
}

func SwitchT(){
    i := 1
    switch i {
        case 2 :
            fmt.Println("2")
            fallthrough
        case 3 :
            fmt.Println(3)
        default :
            fmt.Println("default")
    }
}
