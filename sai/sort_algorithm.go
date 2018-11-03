//排序算法
package main

import "fmt"

func main() {
	// var sli = []string{"a", "b", "c"}
	var sliInt = []int{5, 4, 2, 3, 7}
	// sort1(sli)
	// sort2(sliInt)
	// sort3(sliInt)
	// sort5(sliInt)
	fmt.Println(sliInt)
	quicksort(sliInt, 0, len(sliInt)-1)
	fmt.Println(sliInt)
}

func sort1(sli []string) {

	fmt.Println(sli)
}

//选择
func sort2(sli []int) {
	len := len(sli)
	for i := 0; i < len-1; i++ {
		for j := i + 1; j < len; j++ {
			if sli[i] > sli[j] {
				sli[i], sli[j] = sli[j], sli[i]
			}
		}
	}
	fmt.Println(sli)
}

//冒泡
func sort3(sli []int) {
	len := len(sli)
	for i := 0; i < len-2; i++ {
		for j := i; j < len-1; j++ {
			if sli[i] > sli[j] {
				sli[j], sli[j+1] = sli[j+1], sli[j]
			}
		}
	}
	fmt.Println(sli)
}

//快速排序
func quicksort(sli []int, l, r int) {
	if l < r {
		i, j, x := l, r, sli[l]
		for i < j {

			for i < j && sli[j] >= x {
				j--
			}
			if i < j {
				sli[i] = sli[j]
				i++
			}

			for i < j && sli[i] < x {
				i++
			}
			if i < j {
				sli[j] = sli[i]
				j--
			}
		}

		sli[i] = x
		fmt.Println(sli)
		quicksort(sli, l, i-1)
		quicksort(sli, i+1, r)
	}

}

//快速,不纯正版，适合PHP实现
//TODO 未完成
func sort4(sli []int) {
	len := len(sli)
	if len > 1 {
		mid := sli[0]
		xk, yk := 0, 0
		x := make([]int, len)
		y := make([]int, len)
		for i := 1; i < len; i++ {
			if sli[i] > mid {
				y[xk] = sli[i]
				xk++
			} else {
				x[yk] = sli[i]
				yk++
			}
		}
		sort4(x)
		sort4(y)
		return
	}
}

//快速
func sort5(values []int) {
	if len(values) <= 1 {
		return
	}
	mid, i := values[0], 1
	head, tail := 0, len(values)-1
	for head < tail {
		if values[i] > mid {
			values[i], values[tail] = values[tail], values[i]
			tail--
		} else {
			values[i], values[head] = values[head], values[i]
			head++
			i++
		}
	}
	values[head] = mid
	sort5(values[:head])
	sort5(values[head+1:])
}
