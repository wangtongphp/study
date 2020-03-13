// https://lrita.github.io/2017/12/12/golang-asm/#%E6%B1%87%E7%BC%96%E4%B8%AD%E8%B0%83%E7%94%A8%E5%85%B6%E4%BB%96%E5%87%BD%E6%95%B0
package asm

/*
➜  asm git:(master) ✗ go tool compile -S asm3.1.go
*/

//func caller1(a int) [1]int
func caller2(a int) [1]int
func callee(a int) [1]int
