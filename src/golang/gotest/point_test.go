package gotest

import (
	"fmt"
	"testing"
	"unsafe"
)

//
func Test_Fmt(t *testing.T) {

	a := 1
	v := []*int{&a}

	vv := map[string]*int{"a": &a}

	fmt.Println(v[0], vv["a"])

	fmt.Printf("%v, %#v", v[0], unsafe.Pointer(vv["a"]))

}
