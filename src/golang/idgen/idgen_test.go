package idgen

import (
	"fmt"
	"testing"
)

// go test -race -v golang/idgen -bench . -benchmem -run "Gen"
// Benchmark_GenUID-4         50000             39637 ns/op             256 B/op          6 allocs/op
func Benchmark_GenUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res, _ := genUID()
		fmt.Println(res)
	}
	b.Log("done")
}

// go test -race -v -count=1 golang/idgen -run "Gen"
func Test_GenUID(t *testing.T) {
	res, _ := genUID()
	t.Log(res)
}

//➜  golang git:(master) ✗ go test -race -v golang/idgen -bench . -benchmem -run "Gen" >> idgen/output.log
//➜  golang git:(master) ✗ uniq -c idgen/output.log | wc -l
//   30102
