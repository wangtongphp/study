package calleg

import (
	"testing"
)

//➜  golang git:(master) ✗ go test -v golang/asm/calleg -run "yyy"
func Test_Calleryyy(t *testing.T) {

	y0 := yyy00(11, 22, 33)
	y1 := yyy1(1, 2, 3)
	println("yyy0", y0[0], y0[1], y0[2])
	println("yyy1", y1[0], y1[1], y1[2])

}
