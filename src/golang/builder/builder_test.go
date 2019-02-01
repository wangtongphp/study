/*

wangtong1v:~/devspace/study/src/golang/builder $  /usr/local/go1.11.2.linux-amd64/bin/go test -bench . -benchmem
goos: linux
goarch: amd64
pkg: golang/builder
BenchmarkBuildStringWithBytesBuffer-8         200000        161358 ns/op     1404111 B/op          1 allocs/op
BenchmarkBuildStringWithStringsBuilder-8    100000000           36.8 ns/op        77 B/op          0 allocs/op
PASS
ok      golang/builder  36.267s

*/

//go1.10-examples/stdlib/stringsbuilder/builer_test.go
package builder

import "testing"

func BenchmarkBuildStringWithBytesBuffer(b *testing.B) {
	var builder BuilderByBytesBuffer

	for i := 0; i < b.N; i++ {
		builder.WriteString("Hello, ")
		builder.WriteString("Go")
		builder.WriteString("-1.10")
		_ = builder.String()
	}

}
func BenchmarkBuildStringWithStringsBuilder(b *testing.B) {

	var builder BuilderByStringsBuilder

	for i := 0; i < b.N; i++ {
		builder.WriteString("Hello, ")
		builder.WriteString("Go")
		builder.WriteString("-1.10")
		_ = builder.String()
	}
}
