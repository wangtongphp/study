/**
 * @author wangtong1@xiaomi.com
 */

package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	inByte, _ := ioutil.ReadAll(os.Stdin)
	inStr := string(inByte[0 : len(inByte)-1]) //将最后的\n去掉
	//inStr := string(inByte)
	inSli := strings.Split(inStr, "\n")
	group, _ := strconv.Atoi(inSli[0])
	k := 1
	//遍历组
	for g := 1; g <= group; g++ {
		line, _ := strconv.Atoi(inSli[k])
		k += 1
		//每组数据
		for i := 0; i < line; i++ {

		}

	}

}
