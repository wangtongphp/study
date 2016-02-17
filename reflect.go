package main

import (
    "fmt"
    "reflect"
    )

type Bird struct {
    Name string
    LifeExpectance int
}

func main(){
    sparrow := &Bird{"Sparrow", 3}
    s := reflect.ValueOf(sparrow).Elem()
    typeOfT := s.Type()
    for i := 0; i < s.NumField(); i++ {
        f := s.Field(i)
        fmt.Printf("%d, %s, %s, %v \n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
        fmt.Printf("%d, %s, %s, %v \n", i, reflect.ValueOf(sparrow).Elem().Type().Field(i).Name, reflect.ValueOf(sparrow).Elem().Field(i).Type(), reflect.ValueOf(sparrow).Elem().Field(i).Interface())
    }
}
