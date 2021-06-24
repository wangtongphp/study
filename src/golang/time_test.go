package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func Test_T(t *testing.T) {
	timestamp := time.Now().UnixNano()
	fmt.Println(strconv.FormatInt(timestamp, 10) + "x")
	fmt.Println(2*1e14 + time.Now().UnixNano()/1e6)
}