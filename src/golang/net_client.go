package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	url := "http://localhost:8093/metrics"
	param := map[string]string{
		"a": "1",
		"b": "2",
	}

	c := make(chan os.Signal, 1)
	for {
		signal.Notify(c, syscall.SIGUSR2)

		// Block until a signal is received.
		s := <-c
		for i := 0; i < 2; i++ {
			go func() {
				code, resp, err := GetDataWithQuery(url, param)
				fmt.Println(s, code, string(resp), err)
			}()
		}
	}

}

var httpClient = &http.Client{
	Transport: &http.Transport{
		//DisableKeepAlives: true, //server主动关
		DialContext: (&net.Dialer{
			Timeout:   500 * time.Millisecond,
			KeepAlive: 1 * time.Second,
		}).DialContext,
		MaxIdleConns:          300,
		MaxIdleConnsPerHost:   1,
		IdleConnTimeout:       10 * time.Second,
		ResponseHeaderTimeout: 500 * time.Millisecond,
	},
	Timeout: 800 * time.Millisecond,
}

func GetDataWithQuery(getUrl string, param map[string]string) (statusCode int, response []byte, err error) {

	req, err := http.NewRequest("GET", getUrl, nil)
	if err != nil {
		return 0, nil, err
	}
	if len(param) != 0 {
		q := req.URL.Query()
		for k, v := range param {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, err
	}
	return resp.StatusCode, body, nil
}
