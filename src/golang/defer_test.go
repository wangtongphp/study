/*
实测3倍差距， 书籍中测试5倍差距

wangtong1v:~/devspace/study/src/golang $  go test defer_test.go -bench . -benchmem
BenchmarkDefer-8    50000000            23.3 ns/op         0 B/op          0 allocs/op
PASS
ok      command-line-arguments  1.197s

wangtong1v:~/devspace/study/src/golang $  go test defer_test.go -bench . -benchmem
BenchmarkDefer-8    20000000            67.4 ns/op         0 B/op          0 allocs/op
PASS
ok      command-line-arguments  1.420s


*/
package main

import (
	"sync"
	"testing"
)

var m sync.Mutex

func BenchmarkDefer(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		call()
	}
}

func Test_Defer(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error("recover", err)
		}
	}()
	defer func() {
		t.Log("exec defer")
	}()
	type tmpS struct {
		B int
	}
	var tmpObj struct{ A *tmpS }
	t.Log(tmpObj.A.B)
}
func call() {
	m.Lock()
	// m.Unlock()
	defer m.Unlock()
}
