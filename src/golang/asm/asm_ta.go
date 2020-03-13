// https://lrita.github.io/2017/12/12/golang-asm/#%E6%B1%87%E7%BC%96%E4%B8%AD%E8%B0%83%E7%94%A8%E5%85%B6%E4%BB%96%E5%87%BD%E6%95%B0
package asm

/*
➜  asm git:(master) ✗ go tool compile -S asm3.1.go
"".AsmID SNOPTRDATA size=8
        0x0000 37 25 00 00 00 00 00 00                          7%......
*/

func caller(a int) [1]int
func callerAdv(a int) [1]int
func callee(a int) [1]int

/*
func yyy() {
	ret := zzz(10, 11, 12)
	fmt.Println(ret)
}

func yyy0(a, b, c int) [3]int
func yyy00(a, b, c int) [3]int

//go:noinline
func yyy1(a, b, c int) [3]int {
	return zzz(a, b, c)
}
func zzz(a, b, c int) [3]int

func calleeRel(a, b int) [2]int {
	var d [2]int
	d[0], d[1] = a, b
	return d
}
*/
