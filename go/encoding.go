package main

import(
    //"encoding/base64"
    "fmt"
    "encoding/json"
    "os"
    "github.com/go-simplejson"
)

func main(){
    base64T()
}

func base64T(){
    co := `{"codition":"pay_time > unix_timestamp('2015-07-10 20:21:01' , 'yyyy-MM-dd HH:mm:ss')"}`
    coo, _ := json.Marshal(co)
    js, _ := simplejson.NewJson([]byte(co))
    jss, _ := js.EncodePretty()
    //enc := base64.StdEncoding.EncodeToString(coo)
    fmt.Println("test")
    os.Stdout.Write(coo)
    os.Stdout.Write(jss)

}
