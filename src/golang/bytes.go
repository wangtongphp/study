package main

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

func main() {
	//run()
	bufT()
	//runeT()
	//compareT()
}

//bytes 处理
func compareT() {

	fmt.Println(bytes.Compare([]byte("你好"), []byte("你hao")))

	sl := []byte("s '! ' !")
	sli := []byte("'")
	fmt.Println(string(sl), sli, "\n",
		bytes.Contains(sl, sli), bytes.Count(sl, sli), "\n",
		bytes.Fields(sl), bytes.Fields(sli), "\n",
		bytes.Index(sl, sli), "\n",
		string(bytes.Join([][]byte{sl, sli}, []byte("_______________-----"))), "\n", //php implode
		string(bytes.Replace(sl, []byte("!"), []byte("?"), -1)), "\n",
		bytes.Split(sl, []byte("!")), "\n", //php explode
		bytes.SplitAfter(sl, []byte("!")), "\n", // 前两个切片包含!
	)
}

// 中文的处理
func runeT() {

	s := "hello world, 你好"

	//字节数
	l := len(s)
	fmt.Println("len: ", len(s))

	// 按照字节，错误，中文被分割
	for i := 0; i < l; i++ {
		fmt.Println(i, string(s[i]))
	}

	// 按照字符, 正确做法, 但k遇中文非顺序， k是字节索引非字符索引
	for k, v := range s {
		fmt.Println(k, string(v))
	}

	// 若索引设置到中文字内会出现乱码
	fmt.Println("end", s[13:])

}

// buf
func bufT() {
	buf := bytes.NewBufferString("hello world")
	fmt.Println(buf, buf.String())

	resW, _ := buf.Write([]byte(" []bytes "))
	buf.WriteRune('你')
	buf.WriteByte('!')
	buf.WriteString("好呀")
	fmt.Println(resW, buf.String())

	file, _ := os.Create("/tmp/gocode_t")
	buf.WriteTo(file)

	fmt.Println("WriteTo", buf.String())

	fileR, _ := os.Open("/tmp/gocode_t")
	buf.ReadFrom(fileR)

	fmt.Println("ReadFrom", buf.String())

}

func run() {
	ch := make(chan int, 1024)
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(3e9)
		timeout <- true
	}()

	select {
	case <-timeout:
	case <-ch:
	}

	fmt.Println("d")

}
