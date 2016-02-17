package main

import (
    "fmt"
    )

func sum(values [] int, resChan chan int){
    sum := 0
    for _, val := range values {
        sum += val
    }
    resChan <- sum
}

func main(){
    values := [] int {2,3,4,5,6,3,7,33,55,22}
    resChan := make(chan int, 2)
    go sum(values[:len(values)/2], resChan)
    go sum(values[len(values)/2:], resChan)
    sum1, sum2 := <- resChan, <- resChan
    fmt.Println(resChan, sum1, sum2)
}
    
