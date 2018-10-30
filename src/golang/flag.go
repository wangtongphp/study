package main

import(
    "fmt"
    "strings"
    "flag"
    "os"
//    "time"
)

var flagvar int 

func main(){
    flagT4()
    os.Exit(1)
    flagT3()
    flagT2()
    flagT1()
    flagT()
}

func flagT4(){
    var c string
    flag.StringVar(&c, "c","",`eg: -c='{"1":{"default_range":2,"default_interval":-262},"2":{"default_range":1}}'`)
    flag.Parse()
    fmt.Print(c, flag.Arg(0))
    os.Exit(4)
}

func flagT3(){
    var c string
    flag.StringVar(&c, "c","","eg: -c=1:default_range:2,1:default_interval:-262,2:default_range:1")
    flag.Parse()
    fmt.Print(c)
    co := strings.Split(c,",") 
    row_cus := make(map[string]map[string]interface{})
    //row_cus["a"] = map[string]interface{}{"343":"res"}
    //fmt.Println(row_cus)
    //os.Exit(1)

    for _,v := range co {
        coo := strings.Split(v,":") 
        if len(coo) != 3 {
            fmt.Println("error !=3")
        }
        k0 := coo[0]
        k1 := coo[1]
        row_cus[k0] = map[string]interface{}{k1:coo[2]}

    }
    fmt.Println(row_cus)
}

func flagT2(){
    /*
        visitor := func(f *flag.Flag) {
          if len(f.Name) > 5 && f.Name[0:5] == "test_" {
              g := f.Value.(Getter)
              var ok bool
              switch f.Name {
              case "test_bool":
                  ok = g.Get() == true
              case "test_int":
                  ok = g.Get() == int(1)
              case "test_int64":
                  ok = g.Get() == int64(2)
              case "test_uint":
                  ok = g.Get() == uint(3)
              case "test_uint64":
                  ok = g.Get() == uint64(4)
              case "test_string":
                  ok = g.Get() == "5"
              case "test_float64":
                  ok = g.Get() == float64(6)
              case "test_duration":
                  ok = g.Get() == time.Duration(7)
              }
              if !ok {
                  //t.Errorf("Visit: bad value %T(%v) for %s", g.Get(), g.Get(), f.Name)
              }
          }
      }
      flag.VisitAll(visitor)
      */
}

func flagT(){
    var arg1 = flag.String("arg1","1","input arg1,please")
    flag.IntVar(&flagvar, "arg2", 888, "`arg2` is arg2")
    flag.Parse()
    fmt.Println("g1: ",*arg1)
    fmt.Println("g2: ",flagvar)
    fmt.Println( flag.Args(), flag.Arg(0))
}


type st []string

func(i *st)String() string{
    return fmt.Sprint([]string(*i))
}

func(i *st)Set(value string) error{
    *i = append(*i,value)
    return nil
}

var intervalFlag st 

func flagT1(){

    flag.Var(&intervalFlag, "name", "help message for flagname")
    
    flag.Parse()
    fmt.Println(intervalFlag)

}

