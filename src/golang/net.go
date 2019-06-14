package main

import (
	"fmt"
	"net/http"
)

func main() {
	//http.Handle("/foo", fooHandler)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", r.URL.Path)
	})

	http.ListenAndServe(":8088", nil)
}
