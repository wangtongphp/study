package main

import (
	"context"
	"fmt"
	"strings"
)

func main() {

	r := context.Background()
	ctx := context.WithValue(r, "caller", "xxx")
	c, _ := ctx.Value("traceid").(string)
	fmt.Println(strings.Index(c+"'", "'"))
}
