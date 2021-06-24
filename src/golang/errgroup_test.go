package main

import (
	"errors"
	"golang.org/x/sync/errgroup"
	"testing"
)

//$ go test src/golang/errgroup_test.go -v
func Test_ErrGroup(t *testing.T) {

	t.Log("begin")
	v1:=0
	v2:=0
	var g errgroup.Group
	var err error
	g.Go(func() error {
		v1, err =1,errors.New("err1")
		if err != nil {
			return err
		}
		return nil
	})
	g.Go(func() error {
		v2, err =2,nil
		if err != nil {
			return err
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		t.Log(v1, v2,err)
		return
	}
	t.Log(v1, v2)
	return

}

