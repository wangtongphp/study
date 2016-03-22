package main

import (
    "fmt"
    "os"
   // "strconv"
    "reflect"
    "xmlib/xmjson"
    "encoding/json"
    "io"
    "strings"
    "strconv"
    "github.com/go-simplejson"
)

func main(){
    jsonT()

    simplejsonT()
    os.Exit(1)

    unmarshalBug()
    unmarshalT()
    decoderT()
    xmjsonT()
    xmjsonT1()

}

var jsonDemo = []byte(`[{"a":"ddddd","n":9223372036854775807},{"a":"ssssss","n":9223372036854775807}]`)


//go 解析json　最终解决方案，能准确的读出各种类型的值
func jsonT(){
    js, _:= simplejson.NewJson(jsonDemo)
    vDataHd := js.GetIndex(0).Get("n").Interface()
    switch vDataHd.(type) {
        case json.Number:
            vDataHdi, err := js.GetIndex(0).Get("n").Int64()
            if err != nil {
                vDataHd, _ = js.GetIndex(0).Get("n").Float64()
                vDataHd = strconv.FormatFloat(vDataHd.(float64), 'f', -1, 64)
                fmt.Println("number float64", vDataHd, "\n")
            } else {
                vDataHd = vDataHdi
                fmt.Println("number int64", vDataHd, "\n")
            }
        case string:
            fmt.Println("string", vDataHd, "\n")
    }
    fmt.Println("\n","*************","\n")
}

//用simplejsonT处理9223372036854775807类int64数据
func simplejsonT(){
    js, _:= simplejson.NewJson(jsonDemo)
    d := js.GetIndex(0).Get("n").Interface()
    fmt.Println(d, reflect.TypeOf(d))

    da := js.GetIndex(0).Get("a").Interface()
    fmt.Println(da, reflect.TypeOf(da))

    di,_ := js.GetIndex(1).Get("n").Int64()
    fmt.Println(di, reflect.TypeOf(di))

    ds, _ := js.GetIndex(1).Get("n").String()
    fmt.Println(ds, reflect.TypeOf(ds))

    fmt.Println("\n","*************","\n")
}


//结构体属性必须是大写，json则不是必须
//要想解析在float64到int64之间范围值的必须要提前声明int64，不然会转成float64被截断,官方文档明确注释了
func unmarshalBug(){
    var jst = []byte(`[{"a":"ddddd","n":9223372036854775807},{"a":"ssssss","n":223372036854775807}]`)
    type va struct {
        A string
        N int64
    }

    var vt []va
    err := json.Unmarshal(jst,&vt)
    if err != nil {
        fmt.Println("error:", err)
    }
    fmt.Printf("%+v", vt)
    fmt.Println(vt[1].N, reflect.TypeOf(vt[0].N))

    var r []map[string]interface{}
    err = json.Unmarshal(jst, &r)
    gobook := r[0]
    for k, v := range gobook {
        switch v2 := v.(type) {
            case string:
                fmt.Println(k, "is string", v2)
            case int:
                fmt.Println(k, "is int", v2)
            case float64:
                fmt.Println(k, "is float64", v2)
            case bool:
                fmt.Println(k, "is bool", v2)
            case []interface{}:
                fmt.Println(k, "is an array: ")
                for i, iv := range v2{
                    fmt.Println(i,iv)
                }
            default:
                fmt.Println(k, "is another")
        }
    }
    fmt.Println("\n","*************","\n")

}

//官方案例
func unmarshalT(){
    
    var jsonBlob = []byte(`[
        {"Name": "Platypus", "Order": "Monotremata"},
        {"Name": "Quoll",    "Order": "Dasyuromorphia"}
    ]`)
    type Animal struct {
        Name  string
        Order string
    }
    var animals []Animal
    err := json.Unmarshal(jsonBlob, &animals)
    if err != nil {
        fmt.Println("error:", err)
    }
    fmt.Printf("%+v", animals)
    fmt.Println(animals[0].Name)

}

//Decoder用于解析[多个]json
func decoderT(){

    const jsonStream = `
        {"Name": "Ed", "Text": "Knock knock."}
        {"Name": "Sam", "Text": "Who's there?"}
        {"Name": "Ed", "Text": "Go fmt."}
        {"Name": "Sam", "Text": "Go fmt who?"}
        {"Name": "Ed", "Text": "Go fmt yourself!"}
    `
    type Message struct {
        Name, Text string
    }
    dec := json.NewDecoder(strings.NewReader(jsonStream))
    for {
        var m Message
        if err := dec.Decode(&m); err == io.EOF {
            break
        } else if err != nil {
            fmt.Println("err")
        }
        fmt.Printf("%s: %s\n", m.Name, m.Text)
    }

}

//xmjson处理bigint的bug
func xmjsonT(){
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

    fmt.Println("\n", "********", "\n")
}

//xmjson的用法,(xmjson为内部保密库)
func xmjsonT1(){
    var jstr string = `{
        "test": { 
            "string_slice": ["asdf", "ghjk", "zxcv"],
            "silce": [1, "2", 3],
            "silcewithsubs": [{"subkeyone": 1},
            {"subkeytwo": 2, "subkeythree": 3}],
            "int": 10,
            "float": 5.150,
            "bignum": 9223372036854775807,
            "string": "simplejson",
            "bool": true 
        },
        "test1": { 
            "string_slice": "sb"
        }
    }`

    //xmjson
    js, err := xmjson.NewJSON([]byte(jstr))
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(js)
    
    val, err := js.Get("test1").Get("bignum").String()
    fmt.Println(val, err, "\n")

    silce, err := js.Get("test").Get("silce").Slice()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(silce)

    aws := js.Get("test").Get("silcewithsubs")
    fmt.Println(aws)

    awsval, err := aws.GetIndex(0).Get("subkeyone").Int()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(awsval)

    awsval, err = aws.GetIndex(1).Get("subkeytwo").Int()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(awsval)

    awsval, err = aws.GetIndex(1).Get("subkeythree").Int()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(awsval)

    i, err := js.Get("test").Get("int").Int()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(i)

    f, err := js.Get("test").Get("float").Float64()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(f)

    s, err := js.Get("test").Get("string").String()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(s)

    b, err := js.Get("test").Get("bool").Bool()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(b)

    strs, err := js.Get("test").Get("string_slice").StringSlice()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(strs[0])

    gp, err := js.GetLot("test", "string").String()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(gp)

    gp1, err := js.GetLot("test", "int").Int()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(gp1)

    js.Set("test", "setTest")
    fmt.Println(js.Get("test")) 


    os.Exit(1)

}
