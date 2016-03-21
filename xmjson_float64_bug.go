package main

import (
    "fmt"
    "regexp"
    "time"
    "strconv"
    "xmlib/xmjson"
    )


func main(){

    //xmjson处理bigint的bug
    js, _:= xmjson.NewJSON([]byte(`[9223372036854775807,"dfdf"]`))
    dd, err := js.GetIndex(0).Int64() //-9223372036854775808 <nil>
    df, err := js.GetIndex(0).Float64() //9.223372036854776e+18

    ds:= js.GetIndex(0).Data()
    switch ds.(type) {
        case float64: 
            fmt.Println(strconv.FormatFloat(ds.(float64), 'f', -1, 64))  //9223372036854776000
    }

    fmt.Println(dd, df,err)

    var s float64 = 9223372036854775807
    fmt.Println(strconv.FormatFloat(s,'f',-1,64))
    var si int64 = 9223372036854775807
    fmt.Println(si)


    //字符串转时间戳,类php的strtotime()
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

}
