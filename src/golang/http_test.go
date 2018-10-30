package main

import (
	"fmt"
	"html"
	"net/http"
	"testing"
	"time"
)

func TestClient1(t *testing.T) {

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{
		//CheckRedirect: redirectPolicyFunc,
		Transport: tr,
	}

	resp, err := client.Get("http://example.com")
	t.Log(resp, err)
	// ...

	req, err := http.NewRequest("GET", "http://example.com", nil)
	// ...
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err = client.Do(req)

	t.Log(resp, err)
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func TestServer(t *testing.T) {
	http.Handle("/foo", fooHandler)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	s := &http.Server{
		Addr:           ":80881",
		Handler:        fooHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	t.Log(http.ListenAndServe())
	t.Log(t)
}

func TestServer2(t *testing.T) {

	s := &http.Server{
		Addr:           ":8080",
		Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	t.Log(s.ListenAndServe(":80882", nil))
	t.Log(t)
}
