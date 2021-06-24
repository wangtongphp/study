package main

import (
    "fmt"
    "os"
    //"simplemath"
    "strconv"
)

func Usage(){
    fmt.Println("USAGE: calc command [arguments] ...","The commands are:","sqrt 开方,eg: calc sqrt 4 => 2", "add  加和,eg: calc add 2 2 => 4")
}

func main(){
    args := os.Args[1:]
    opt := args[0]
    switch opt {
        case "add" : 
            r1, err1 := strconv.Atoi(args[1])
            r2, err2 := strconv.Atoi(args[2])
            if(err1 != nil || err2 != nil){
                fmt.Println("add strconv error", r1, r2)
            }
            //res := simplemath.Add(r1, r2)
            //fmt.Println(res)
/*
        case "sqrt" :
            r1, err1 := strconv.Atoi(arg[1])
            if(err1 != nil){
                fmt.Println("sqrt strconv err", err1)
            }
            res := simplemath.sqrt(r1)
            fmt.Println(res)
*/
        default :
            Usage()

    }
}
