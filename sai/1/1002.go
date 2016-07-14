/*
 * @author wangtong1@xiaomi.com
 * @desc 此题注意，每行输出的最后一位不能有空格
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
		l := strings.Trim(input[2+i*2], " ")
		tmp := strings.Split(l, " ")
		run(tmp)
	}
}

func run(sli []string) {
	//fmt.Print(sli)

	len := len(sli)
	res := make([]string, 0, len)
	var a int = 0
	j := 3
	for {
		if a >= len-1 {
			res = append(res, sli[len-1])
			break
		}
		res = append(res, sli[a])
		a = a + j
		j = j + 2
	}

	fmt.Printf("%s", strings.Join(res, " "))
	fmt.Print("\n")
}

/**
神奇数字
题目描述：
某天你突然得到一串秘密数字，M 告诉你解密的方法是，将这些数字按照第
一行 1 个，第二行 3 个，第三行 5 个，第 n 行(n-1)*2+1 个的顺序排放这些数，
然后将每行最后一个数组合起来就是最终的密码。
你需要完成这个任务。
示例：
1 2 3 4 5 6 7 8 9 10
摆放:
1
2 3 4
5 6 7 8 9
10
则密码为：1 4   9   10
输入：
输入第一行为 T，表示有 T 组测试数据， （1<=T<=10）。
接下来有 T 组测试数据，每组数据第一行是一个数 n (   1<=n    <=  1000)，表示
有 n 个神秘数字，接下来一行，包含 n 个整数 a0,  a1  …   an-1，相邻之间用一个空
格分开，对每个数字 ai ( 0   <=  ai  <=  1000,   0   <=  int<    n)。
输出：
你的è¾出包含 T 行，每行是一组测试数据的结果，
对每组测试数据，输出若干个数字，也即密码，数字之间用空格分开。
样例输入：
3
10
1 2 3 4 5 6 7 8 9 10
2
2 1
6
4 3 2 2 5 2
样例输出：
1   4   9   10
2   1
4   2   2

*/
