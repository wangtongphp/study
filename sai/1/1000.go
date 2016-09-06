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

func main() {
	inputB, _ := ioutil.ReadAll(os.Stdin)
	inputS := string(inputB[0 : len(inputB)-1])
	input := strings.Split(inputS, "\n")

	line, _ := strconv.Atoi(input[0])
	for i := 0; i < line; i++ {
		//n, _ := strconv.Atoi(input[1+i*2])
		//tmp := make([]string, 0, n)
		l := strings.Trim(input[i+1], " ")
		tmp := strings.Split(l, " ")
		run(tmp)
	}

}
func run(sli []string) {
	nA, _ := strconv.Atoi(sli[0])
	nB, _ := strconv.Atoi(sli[1])
	fmt.Printf("%v\n", nA+nB)
}

/**
A. a+b
题目描述：
作为热身题，题目也相当简单：给出两个正整数 a 和 b，输出 a+b 的结果。
输入：
输入第一行为 T，表示测试数据组数。
接下来有 T 组测试数据，每组数据一行，每行有两个正整数， a 和 b，空格
分开。(a 和 b 保证小于 2^32)
输出：
你的输出包含 T 行，每行是一组测试数据的结果，也即 a+b 的结果。
样例输入：
3
1   2
3   4
5   6
样例输出：
3
7
11
*/
