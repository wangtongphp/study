package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	//httpT()
	webT()
}

func webT() {
	http.HandleFunc("/h", hellowHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

func hellowHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world")
}

func httpT() {
	req, err := http.NewRequest("GET", "http://www.mi.com", nil)
	req.Header.Add("User-Agent", "wangtong-PC")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	io.Copy(os.Stdout, resp.Body)
	fmt.Println(resp)
}
