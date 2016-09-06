package simplemath

import (
    "testing"
    "fmt"
    )

func TestAdd1(){
    res := add(3, 5)
    if(res != 8){
        fmt.Println("testing err: 3+5=",res)
        t.Errorf("Add(3,5) failed. Got %d, expected 3.",res)
    }
}

func TestAdd2(){
    res := add(8, 5)
    if(res != 13){
        fmt.Println("testing err: 8+5=",res)
        t.Errorf("Add(8,5) failed. Got %d, expected 3.",res)
    }
}
