package asm

import (
	"testing"
)

//➜  golang git:(master) ✗ go test -v golang/asm
//=== RUN   Test_Caller
//[10 11 12]
func Test_CallerOne(t *testing.T) {

	c := caller2(10)
	println("c", c[0])

}
