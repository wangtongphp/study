package main

var helloworld = "你好, 世界"

func main() {
	//println(asm.Id, asm.Name)
	Foo()
}

// ➜  golang git:(master) ✗ go build -gcflags "-N -l" -o main main.go;  go tool objdump -s "main.Foo" main
func Foo() {
	var c []byte
	var b int16
	var a bool

	_, _, _ = a, b, c
}
