package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
	//"io/ioutil"
)

func main() {
	//slAppend()
	//slForeach()
	//var sl_t = []interface{}{"0",8,22,11}
	//var res = slIn(0, sl_t)
	//res := slRand(sl_t)
	//fmt.Println(res)
	//timeF()
	//stringSplit()
	//templateMain()
	//fileMain()
	//moreSli()
	//sl2()
	//NilSli()

	//<-make(chan int,0)
	time.Sleep(time.Second*8)
}

func NilSli(){
	var a []int
	var b = make([]int,0)
	fmt.Println(a==nil, b==nil,a,b)
	for i,v:=range a{
		fmt.Println(i,v)
	}
	for i,v:=range b{
		fmt.Println(i,v)
	}
	fmt.Println(b[0])
}

// .
func sl2() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover....", err)
		}
	}()
	defer func(){
		fmt.Println()
	}()
	val:= []int64{
		0,1,2,3,
	}
	v:=val[len(val)-1]
	fmt.Println(v)
}



// slice 不能用索引赋值
func sl1() {
	val:= make(map[int64][]int64)
	for i:=0;i<3;i++{
		val[1] = append(val[1],int64(i))
	}
	fmt.Println(val)
}

// slice 不能用索引赋值
func moreSli() {
	type item struct{ ID int }
	type sliVal []*item
	a := make([]sliVal, 0, 220101)

	aa := sliVal{&item{11}}
	//a[1] = append(a[1], &item{33})
	fmt.Println(aa, a)

}

func fileMain() {
	f, err := os.OpenFile("/tmp/tmp.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	f.WriteString("heelooo")
	fmt.Println(FileExists("tmp.txt"), SelfDir(), SelfPath())
}

func FileExists(name string) (is bool) {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func SelfDir() (dir string) {
	dir = filepath.Dir(SelfPath())
	return
}

func SelfPath() (path string) {
	path, _ = filepath.Abs(os.Args[0])
	return
}

func templateMain() {
	ss := make([]*School, 0)
	for i := 0; i < 4; i++ {
		s := new(School)
		s.Group = i
		s.Student["age"] = fmt.Sprintf("%d", 20+i)
		s.Student["name"] = fmt.Sprintf("name%d", 1+i)
		ss = append(ss, s)
	}
	fmt.Println(ss)

	t := template.New("test").Funcs(tplFuncMap)
	t, _ = t.Parse(`Xiaomi :
        {{range $k, $v := .}}
        group : {{$v.Group}}
        {{if eq $v.Group 2}} ifyes {{end}}
        student info: 
        {{range $kk, $vv :=$v.Student}}
        {{$kk}} is {{$vv}}
        {{end}}
        {{end}}`)
	t.Execute(os.Stdout, ss)

}

var (
	tplFuncMap template.FuncMap = make(template.FuncMap)
)

func init() {
	tplFuncMap["eq"] = Equal
}

func Equal(a, b interface{}) (res bool) {
	res = a == b
	return
}

type School struct {
	Group   int
	Student map[string]string
}

func stringSplit() {
	var str string = "xiaomi hello'"
	strSl := strings.Split(str, "o")
	str1 := strings.Index(str, "o")
	str2 := strings.LastIndex(str, "o")
	fmt.Println(str[:6], strSl, strSl[0], str[str1:], str2)
}

func timeF() {
	var stamp int64 = time.Now().Unix()
	now := time.Unix(stamp, 0)
	dt := time.Now().Format("2006-01-02 15:04:05")
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	date, err := time.ParseInLocation("20060102", "20150302", loc)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(now, dt, date)
}

/*
func slRand(sl []interface{})(res interface{}){
    randnum := rand.Intn(len(sl))
    res = sl[randnum]
    return
}
*/

func slAppend() {
	sl := make([]string, 0)
	sl = append(sl, "1")
	sl = append(sl, "4")
	sl = append(sl, "8")
	fmt.Println(sl)
}

func slForeach() {
	var sl = []interface{}{"22", "44"}
	fmt.Println(sl)
	//var sl8 = []interface{}{}
	sl8 := make([]interface{}, 0)
	for k, v := range sl {
		sl8 = append(sl8, v)
		fmt.Println(k, v)
	}
	fmt.Println(sl8)
}

func slIn(i interface{}, sl8 []interface{}) (b bool) {
	for _, v := range sl8 {
		if v == i {
			return true
		}
	}
	return false
}
