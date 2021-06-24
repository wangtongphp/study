package main

import "fmt"

func main(){
	a:= []bool{
		false,
		false,
		false,
		true,
		false,
		true,
		false,
		false,
		false,
		true,
		false,
	}
	sliceHandler(a)
	fmt.Println(a)
}

//filter header&tail false.
//midden change more zero to single zero
func sliceHandler(sli []bool) {
	resPoint :=0
	var prev bool //上一个值
	var stage bool //是否进入新策略
	for _,v:=range sli{
		if !stage  && !v{
			prev = v
			continue
		} else if !stage && v{
			resPoint++
			sli[resPoint] = true
			prev = v
			stage = true
			continue
		}else if stage && prev{
			resPoint++
			sli[resPoint] = false
			prev = v
		}else if stage && !prev{
			prev = v
			continue
		}
	}
	sli = sli[0:resPoint]
	return
}