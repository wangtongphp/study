package main

import(
    "time" 
    "fmt"
    "bytes"
    "os"
    )

func main(){
    //run()
    //bufT()
    //runeT()
    compareT()
}

func compareT(){

    fmt.Println(bytes.Compare([]byte("你好"), []byte("你hao")))

    sl := []byte("s '! ' !")
    sli := []byte("'")
    fmt.Println(string(sl), sli, 
        bytes.Contains(sl,sli), bytes.Count(sl,sli),
        bytes.Fields(sl), bytes.Fields(sli),
        bytes.Index(sl, sli),
        bytes.Join([][]byte{sl,sli}, []byte("_______________-----")), //php implode
        string(bytes.Replace(sl,[]byte("!"), []byte("?"), -1)),
        bytes.Split(sl,[]byte("!")), //php explode
        bytes.SplitAfter(sl,[]byte("!")), // 前两个切片包含!
        )
}

func runeT(){

    s := "hello world, 你好"

    l := len(s)

    for i:=0; i<l; i++{
        fmt.Println(i, s[i])
    }

    for  k,v := range s{

        fmt.Println(k, v)
    }

    fmt.Println(s[4:5])

}


func bufT(){
    buf := bytes.NewBufferString("hello world")
    fmt.Println(buf, buf.String())

    resW,_ := buf.Write([]byte(" []bytes "))
    buf.WriteRune('你')
    buf.WriteByte('!')
    buf.WriteString("好呀")
    fmt.Println(resW, buf.String())

    file,_ := os.Create("/tmp/gocode_t")
    buf.WriteTo(file)
   
    fmt.Println("WriteTo", buf.String())

    fileR,_ := os.Open("/tmp/gocode_t")
    buf.ReadFrom(fileR)


    fmt.Println("ReadFrom", buf.String())


}

func run(){
    ch := make(chan int, 1024)
    timeout := make(chan bool,1)
    go func() {
        time.Sleep(3e9)
        timeout <- true
    }()

    select{
        case <-timeout:
        case <-ch:
    }

    fmt.Println("d" )

}




