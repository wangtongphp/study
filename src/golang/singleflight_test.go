package main


import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"testing"
	"time"

	//"time"
)

var sf singleflight.Group

func Test_FFF(t *testing.T) {

	for i := 0; i < 99; i++ {
		time.Sleep(time.Second*1)
		go func() {
		v, err, s := sf.Do("key", func() (i interface{}, err error) {
			time.Sleep(time.Second*3)
			fmt.Println("sf.Do")
			return "svc.GetExamDetailByQSID(ctx, qsID)", err
		})
		fmt.Println(v, err, s)
	}()
	}

/*
	for i := 0; i < 9; i++ {
			v, err, s := sf.Do("key", func() (i interface{}, err error) {
				fmt.Println("sf.Do")
				return "svc.GetExamDetailByQSID(ctx, qsID)", err
			})
			fmt.Println(v, err, s)
	}
*/
	select{}
}
