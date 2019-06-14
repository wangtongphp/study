package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/signal"
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
		signal.Notify(c, os.Interrupt)

		// Block until a signal is received.
		s := <-c
		code, resp, err := GetDataWithQuery(url, param)
		fmt.Println(s, code, string(resp), err)
	}

}

func GetDataWithQuery(getUrl string, param map[string]string) (statusCode int, response []byte, err error) {

	var httpClient = &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   500 * time.Millisecond,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:          300,
			MaxIdleConnsPerHost:   100,
			IdleConnTimeout:       30 * time.Second,
			ResponseHeaderTimeout: 500 * time.Millisecond,
		},
		Timeout: 800 * time.Millisecond,
	}

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
