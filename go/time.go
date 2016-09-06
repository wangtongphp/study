package main

import (
    "fmt"
    "regexp"
    "time"
    "strconv"
    )


func main(){
    //strtotimeT()
    intToDt()
}

//字符串转时间戳,类php的strtotime()
func strtotimeT(){
    var v string = "3"
    vDataHd := "2015-07-11 19:47:05"
    if matched, _ := regexp.MatchString(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}$`, vDataHd); matched {
        loc, _ := time.LoadLocation("Asia/Shanghai")
        t, _ := time.ParseInLocation("2006-01-02 15:04:05", vDataHd, loc)
        vv:= t.Unix()
        v = strconv.FormatInt(vv,10)
    }

    fmt.Println(v)
    //time_u := 1436644025

    fmt.Println("\n", "********", "\n")
}

//int 转datetime
func intToDt(){
    var v string = "1409988776"
    vi, err := strconv.ParseInt(v, 10, 64)
    t := time.Unix(vi,0)
    tt := t.String()
    dt := t.Format("2006-01-02 15:04:05")
    fmt.Println(tt,dt, err)
}
