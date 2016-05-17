package main

import(
    //"io"
    "os"
    "io/ioutil"
    "fmt"
    "strconv"
    )

func main(){
    fmt.Println("hello")
    run(10)    
}

func run(max int){
     
    bytes, err := ioutil.ReadFile("i1.txt")
    os.Stdout.Write(bytes)
    //err = ioutil.WriteFile("i1.txt",[]byte(`write,write`), 0777)

    file, err := os.Open("i1.txt")
    data := make([]byte,100)
    cnt, err := file.Read(data)
    fmt.Printf("read %d byte: %q\n", cnt, data[:cnt])

    slice := append([]string{"wwwwwwww"}, "ooo")
    //os.Stdout.Write(slice)
    fmt.Println(slice)



    if err != nil {
        fmt.Println(err)
    }
    var a=1
    resA := "1"
    j := 1
    for {
       if(a > max){
            resA = resA+strconv.Itoa(max)
            break;
       }
       j = j+2
       a = a + j
       resA = resA+strconv.Itoa(a)
    }
        
    fmt.Println(resA)
}


