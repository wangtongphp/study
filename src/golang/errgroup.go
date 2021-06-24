package main

import (
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
)

func main() {
	fmt.Println("begin")
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
		v2, err =2,errors.New("err2")
		if err != nil {
			return err
		}
		return nil
	})
	g.Go(func() error {
		defer func(){
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		v2, err =3,nil
		if err != nil {
			return err
		}
		err.Error()
		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Println(v1,v2,err)
		return
	}
	fmt.Println(v1,v2,err)
	return

}

