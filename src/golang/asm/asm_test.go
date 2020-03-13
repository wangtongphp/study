package asm

import (
	"testing"
)

/*
//➜  golang git:(master) ✗ go test -v golang/asm
//=== RUN   Test_Caller
//[10 11 12]
func Test_Calleryyy(t *testing.T) {

	y0 := yyy00(1, 2, 3)
	y1 := yyy1(1, 2, 3)
	println("yyy0", y0[0], y0[1], y0[2])
	println("yyy1", y1[0], y1[1], y1[2])

}
*/

//➜  golang git:(master) ✗ go test -v golang/asm
//=== RUN   Test_Caller
//[10 11 12]
func Test_CallerOne(t *testing.T) {

	c := callerAdv(11)
	println("callee", c[0])

}
