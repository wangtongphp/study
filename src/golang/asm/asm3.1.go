package asm

/*
➜  asm git:(master) ✗ go tool compile -S asm3.1.go
"".AsmID SNOPTRDATA size=8
        0x0000 37 25 00 00 00 00 00 00                          7%......
*/
//var AsmID = 9527 // 0x2537

//var asmName = "abcdefghijk"

var Id int

var NameData [8]byte
var Name string
