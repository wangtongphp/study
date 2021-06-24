package main

import(
    "log"
    "bytes"
    "fmt"
    "os"
    "time"
    "github.com/sirupsen/logrus"
)

func init(){
    defer func(){
        fmt.Println("init defer")
    }()
}

func main(){
    defer func(){
        fmt.Println("main defer")
    }()
    logrusT()
    logT()
    os.Exit(1)
}


func logrusT(){

    logrus.SetLevel(logrus.DebugLevel)
    logr := logrus.New()
    logr.Debug("debuggg")
    logr.Info("this is info")
    time.Sleep(1e9)
    logr.Warn("warning")
    logr.Error("errorrr")

    logrus.SetFormatter(&logrus.JSONFormatter{})
    logrus.Info("JSONFormatter")

    logrus.SetFormatter(&logrus.TextFormatter{})
    logrus.Info("TextFormatter")


    //logr.Fatal(logrus.GetLevel())
}

func logT(){
    logger := log.New(os.Stdout, "logger: ", log.Lshortfile)
    logger.Print("Hello, log file!")
    time.Sleep(1e9)
    logger.Print("\033[1;4;32m Hello, logiiiii file! \033[0m")

    var buf bytes.Buffer
    logger = log.New(&buf, "logger: ", log.Lshortfile)
    logger.Print("Hello, log file!")
    time.Sleep(1e9)
    logger.Print("Hello, logiiiii file!")
    fmt.Print(&buf)

}


