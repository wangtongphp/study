package gotest

import (
	"strings"
	"testing"
)

// go test -v golang/gotest -run "MapKey"
func Test_MapKey(t *testing.T) {

	a := "a:"
	m := make(map[string]int, 0)
	m[a] = 1
	t.Log(m[a])

	s := strings.Split(a, ":")
	t.Log(len(s))
}
