// https://colobu.com/2017/06/21/json-tricks-in-Go/
// 如果 PHP array是空的时候，序列化出来是[]
//  A int `json:"a,string,omitempty"` 无论在Marshal或UnMarshal都可以互转string和int，转失败会报错值为0
//  如何优雅的实现，判断解析出来为0的值是真的0还是不存在？

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	//"os"
	//// "strconv"
	//"bytes"
	//"encoding/gob"
	//"encoding/json"
	//"github.com/bitly/go-simplejson"
	//"github.com/json-iterator/go"
	//"github.com/json-iterator/go/extra"
	//"io"
	//"reflect"
	//"strconv"
	//"strings"
	//"xmlib/xmjson"
)

func main() {
	j3()
	//j2()
	// j14()
	//j7()
	// jsonPHPEncode()
	// jsonEncode()
	// jsonEncodeT()
	// os.Exit(1)
	//
	// jsonT()
	//
	// simplejsonT()
	// os.Exit(1)
	//
	// unmarshalBug()
	// unmarshalT()
	// decoderT()
	// xmjsonT()

}

func j3() error{
	err:= errors.New("sdf")
	defer func(){
		fmt.Println( err)
	}()
	err = errors.New("aaa")
	return errors.New("bbb")
}


func j2() {
	//v:=make([]int,0)
	var v []int
	vv,e:=json.Marshal(v)
	fmt.Printf("%+v,%+v,%+v",vv,string(vv),e)
}

//
//func j14() {
//
//	extra.SetNamingStrategy(extra.LowerCaseWithUnderscores)
//	output, _ := jsoniter.Marshal(struct {
//		UserName      string `json:"-"`
//		FirstLanguage string //`json:"first_language"`
//		Age           int    `json:",string,omitempty"`
//	}{
//		UserName:      "taowen",
//		FirstLanguage: "Chinese",
//		Age:           0,
//	})
//	fmt.Printf("%s, %b \n", output, (`{"user_name":"taowen","first_language":"Chinese"}` == string(output)))
//
//}
//
//func j7() {
//
//	//PHP应该规范用：  json_encode($foo,JSON_FORCE_OBJECT), 或者 (object)array() ， new stdClass()
//	extra.RegisterFuzzyDecoders() //这个能避免[]类型对应错误导致的报错
//	var d1 = []byte(`{"a":"32","b":[]}`)
//	//var d1 = []byte(`{"a":"32","b":{"c":0}}`) // 这个和上面json不一样，但Unmarshal却相同
//	type bs struct {
//		C int `json:"c,omitempty"`
//	}
//	type ts struct {
//		A int    `json:"a,string,omitempty"`
//		B bs     `json:"b,omitempty"`
//		D string `json:"d,omitempty"`
//	}
//	var res ts
//	err := jsoniter.Unmarshal(d1, &res)
//	if err != nil {
//		fmt.Println(err, res)
//	}
//	/**
//	  j7(): {A:32 B:{C:0} D:}
//	  0, int, %!v(MISSING)
//	*/
//	fmt.Printf("j7(): %+v\n%#v, %T, %#v\n", res, res.B.C, res.B.C)
//
//}
//
////TODO 如何优雅的实现，判断解析出来为0的值是真的0还是不存在？
//func jsonPHPEncode() {
//	// php -r 'echo json_encode(array("a"=>3,"b"=>array()));';
//	// {"a":3,"b":[]}
//	var d1 = []byte(`{"a":32,"b":[]}`)
//	// var d1 = []byte(`{"b":3}`)
//	// var d1 = []byte(`{"a":3,"b":{"c":4}}`)
//	type bs struct {
//		C int `json:"c,omitempty"`
//	}
//	type ts struct {
//		A int `json:"a,omitempty"`
//		B bs  `json:"b,omitempty"`
//	}
//	var res ts
//	err := json.Unmarshal(d1, &res)
//	if err != nil {
//		fmt.Println(err, res)
//	}
//	/**
//	  {A:3 B:{C:0}}
//	  0, int
//	*/
//	fmt.Printf("jsonPHPEncode(): %+v\n%#v, %T\n", res, res.B.C, res.B.C)
//
//}
//
//// 注意，要加omitempty, 防止数据原义发生变化
//func jsonEncode() {
//	type bs struct {
//		C int `json:"c"`
//	}
//	type ts struct {
//		A int `json:"a,omitempty"`
//		B bs  `json:"b,omitempty"`
//	}
//	orig := ts{
//		A: 0, //若为0值，omitempty了，将会忽略此字段， 坑爹...
//		B: bs{
//			// C: 4, //若没有这行，也没哟omitempty，则会解析为 {"a":3,"b":{"c":0}}
//		},
//	}
//	res, err := json.Marshal(orig)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Printf("%+s\n", res)
//
//}
//
////go 解析json　最终解决方案，能准确的读出各种类型的值
//func jsonT() {
//	var jsonDemo = []byte(`[{"a":"ddddd","n":9223372036854775807},{"a":"ssssss","n":9223372036854775807}]`)
//	js, _ := simplejson.NewJson(jsonDemo)
//	vDataHd := js.GetIndex(0).Get("n").Interface()
//	switch vDataHd.(type) {
//	case json.Number:
//		vDataHdi, err := js.GetIndex(0).Get("n").Int64()
//		if err != nil {
//			vDataHd, _ = js.GetIndex(0).Get("n").Float64()
//			vDataHd = strconv.FormatFloat(vDataHd.(float64), 'f', -1, 64)
//			fmt.Println("number float64", vDataHd, "\n")
//		} else {
//			vDataHd = vDataHdi
//			fmt.Println("number int64", vDataHd, "\n")
//		}
//	case string:
//		fmt.Println("string", vDataHd, "\n")
//	}
//	fmt.Println("\n", "*************", "\n")
//}
//
////"{\"codition\":\"pay_time > unix_timestamp('2015-07-10 20:21:01' , 'yyyy-MM-dd HH:mm:ss')\"}"
//func jsonEncodeT() {
//	co := `{"codition":"pay_time > unix_timestamp('2015-07-10 20:21:01' , 'yyyy-MM-dd HH:mm:ss')"}`
//	cio := map[string]interface{}{"table_name": "xm_stat_alipay", "codition": "pay_time > unix_timestamp('2015-07-10 20:21:01' , 'yyyy-MM-dd HH:mm:ss')"}
//
//	coo, _ := json.Marshal(cio)
//	os.Stdout.Write(coo)
//
//	js, _ := simplejson.NewJson([]byte(coo))
//	jss, _ := js.EncodePretty()
//	os.Stdout.Write(jss)
//
//	var buf bytes.Buffer
//	enc := gob.NewEncoder(&buf)
//	err := enc.Encode(cio)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println("\n", "buf.Bytes()")
//	os.Stdout.Write(buf.Bytes())
//
//	realJ, _ := json.Marshal(cio)
//	fmt.Println("\n", "realJ")
//	os.Stdout.Write(realJ)
//
//	//cii := make([]map[string]interface{},`{"codition":"p"}`,`{"table":"xm"}`)
//	//cij, _ := json.Marshal(cii)
//	//fmt.Println(cij)
//
//	type cst struct {
//		t interface{}
//		c interface{}
//	}
//	cs := &cst{}
//	cs.t = cio["table_name"]
//	cs.c = cio["codition"]
//	fmt.Println("\n struct ", cs)
//	csj, _ := json.Marshal(cs)
//	fmt.Println("\n struct encode", csj)
//
//	var realSs string
//	if err := json.Unmarshal([]byte(co), &realSs); err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(realSs)
//
//}
//
////用simplejsonT处理9223372036854775807类int64数据
//func simplejsonT() {
//	var jsonDemo = []byte(`[{"a":"ddddd","n":9223372036854775807},{"a":"ssssss","n":9223372036854775807}]`)
//	js, _ := simplejson.NewJson(jsonDemo)
//	d := js.GetIndex(0).Get("n").Interface()
//	fmt.Println(d, reflect.TypeOf(d))
//
//	da := js.GetIndex(0).Get("a").Interface()
//	fmt.Println(da, reflect.TypeOf(da))
//
//	di, _ := js.GetIndex(1).Get("n").Int64()
//	fmt.Println(di, reflect.TypeOf(di))
//
//	ds, _ := js.GetIndex(1).Get("n").String()
//	fmt.Println(ds, reflect.TypeOf(ds))
//
//	fmt.Println("\n", "*************", "\n")
//}
//
////结构体属性必须是大写，json则不是必须
////要想解析在float64到int64之间范围值的必须要提前声明int64，不然会转成float64被截断,官方文档明确注释了
//func unmarshalBug() {
//	var jst = []byte(`[{"a":"ddddd","n":9223372036854775807},{"a":"ssssss","n":223372036854775807}]`)
//	type va struct {
//		A string
//		N int64
//	}
//
//	var vt []va
//	err := json.Unmarshal(jst, &vt)
//	if err != nil {
//		fmt.Println("error:", err)
//	}
//	fmt.Printf("%+v", vt)
//	fmt.Println(vt[1].N, reflect.TypeOf(vt[0].N))
//
//	var r []map[string]interface{}
//	err = json.Unmarshal(jst, &r)
//	gobook := r[0]
//	for k, v := range gobook {
//		switch v2 := v.(type) {
//		case string:
//			fmt.Println(k, "is string", v2)
//		case int:
//			fmt.Println(k, "is int", v2)
//		case float64:
//			fmt.Println(k, "is float64", v2)
//		case bool:
//			fmt.Println(k, "is bool", v2)
//		case []interface{}:
//			fmt.Println(k, "is an array: ")
//			for i, iv := range v2 {
//				fmt.Println(i, iv)
//			}
//		default:
//			fmt.Println(k, "is another")
//		}
//	}
//	fmt.Println("\n", "*************", "\n")
//
//}
//
////官方案例
//func unmarshalT() {
//
//	var jsonBlob = []byte(`[
//        {"Name": "Platypus", "Order": "Monotremata"},
//        {"Name": "Quoll",    "Order": "Dasyuromorphia"}
//    ]`)
//	type Animal struct {
//		Name  string
//		Order string
//	}
//	var animals []Animal
//	err := json.Unmarshal(jsonBlob, &animals)
//	if err != nil {
//		fmt.Println("error:", err)
//	}
//	fmt.Printf("%+v", animals)
//	fmt.Println(animals[0].Name)
//
//}
//
////Decoder用于解析[多个]json
//func decoderT() {
//
//	const jsonStream = `
//        {"Name": "Ed", "Text": "Knock knock."}
//        {"Name": "Sam", "Text": "Who's there?"}
//        {"Name": "Ed", "Text": "Go fmt."}
//        {"Name": "Sam", "Text": "Go fmt who?"}
//        {"Name": "Ed", "Text": "Go fmt yourself!"}
//    `
//	type Message struct {
//		Name, Text string
//	}
//	dec := json.NewDecoder(strings.NewReader(jsonStream))
//	for {
//		var m Message
//		if err := dec.Decode(&m); err == io.EOF {
//			break
//		} else if err != nil {
//			fmt.Println("err")
//		}
//		fmt.Printf("%s: %s\n", m.Name, m.Text)
//	}
//
//}


/*
//xmjson处理bigint的bug
func xmjsonT() {
	js, _ := xmjson.NewJSON([]byte(`[9223372036854775807,"dfdf"]`))
	dd, err := js.GetIndex(0).Int64()   //-9223372036854775808 <nil>
	df, err := js.GetIndex(0).Float64() //9.223372036854776e+18

	ds := js.GetIndex(0).Data()
	switch ds.(type) {
	case float64:
		fmt.Println(strconv.FormatFloat(ds.(float64), 'f', -1, 64)) //9223372036854776000
	}

	fmt.Println(dd, df, err)

	var s float64 = 9223372036854775807
	fmt.Println(strconv.FormatFloat(s, 'f', -1, 64))
	var si int64 = 9223372036854775807
	fmt.Println(si)

	fmt.Println("\n", "********", "\n")
}
*/
