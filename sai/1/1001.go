/**
 * @status 通过
 */

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var sli = make([]int, 50)

func main() {
	inputB, _ := ioutil.ReadAll(os.Stdin)
	inputS := string(inputB[0 : len(inputB)-1])
	input := strings.Split(inputS, "\n")

	line, _ := strconv.Atoi(input[0])
	for i := 0; i < line; i++ {
		//n, _ := strconv.Atoi(input[1+i*2])
		//tmp := make([]string, 0, n)
		nS := strings.Trim(input[i+1], " ")
		nI, _ := strconv.Atoi(nS)
		n := run(nI)
		fmt.Printf("%v\n", n)
		//fmt.Print(sli)
	}

}
func run(nI int) (n int) {
	if nI <= 0 {
		return 0
	} else if nI == 1 {
		return 1
	} else if nI == 2 {
		return 2
	} else {
		//cache
		if sli[nI] == 0 {
			n := run(nI-1) + run(nI-2)
			sli[nI] = n
			return n
		}
		return sli[nI]
	}
}

/**
上楼梯
题目描述：
总参三层有个楼梯，一共有 n 级，你从地板开始，每次可以上一级，或者上
两级，问走到第 n 级一共有多少种走法？
输入：
第一行为正整数 t，表示有多少组测试数据。
接下来的 t 行，每行一个正整数 n，(1<=n<=50)，表示楼梯有 n 级。
输出：
输出一共 t 行，每行一个整数，表示走 n 级有多少种方案。
样例输入：
3
1
2
5
样例输出：
1
2
8
*/
