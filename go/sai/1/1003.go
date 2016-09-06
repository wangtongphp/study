/**
 * @TODO 运行超时
 * @author wangtong1
 */

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
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
	for i := 1; i <= group; i++ {
		provI, _ := strconv.Atoi(inSli[k])
		k += 1

		//结果，默认0
		provRes := make(map[int]int)

		//省配置
		provMap := make(map[int]map[string]int)
		for p := 1; p <= provI; p++ {
			provSli := strings.Split(inSli[k], " ")
			provKey, _ := strconv.Atoi(provSli[2])
			provMap[provKey] = map[string]int{"min": ip2long(provSli[0]), "max": ip2long(provSli[1])}
			provRes[provKey] = 0
			k += 1
		}
		//fmt.Println(provMap)
		ipI, _ := strconv.Atoi(inSli[k])
		k += 1

		//ip
		for m := 1; m <= ipI; m++ {
			long := ip2long(inSli[k])
			k += 1
			for pk, pv := range provMap {
				if long >= pv["min"] && long <= pv["max"] {
					provRes[pk] += 1

				}
			}
		}

		//sort
		var resSli []int
		for k, _ := range provRes {
			resSli = append(resSli, k)
		}
		sort.Ints(resSli)

		//print
		for _, key := range resSli {
			fmt.Printf("%v %v\n", key, provRes[key])
		}
	}

}

func ip2long(ip string) (long int) {
	ipSli := strings.Split(ip, ".")
	ip1, _ := strconv.Atoi(ipSli[0])
	ip2, _ := strconv.Atoi(ipSli[1])
	ip3, _ := strconv.Atoi(ipSli[2])
	ip4, _ := strconv.Atoi(ipSli[3])
	ip1 = ip1 << 24
	ip2 = ip2 << 16
	ip3 = ip3 << 8
	ip4 = ip4
	long = ip1 + ip2 + ip3 + ip4
	return long
}

/**
IP 地址统计
题目᧿述：
数据中心有一些 ip 地址库，包括每个地址区间以及归属省份(ip 地址区间不
会重叠交叉)，现在要对用户访问日志统计每个省份的访问数量。
输入：
输入第一行为 T，表示包含 T 组测试数据。
对于每组测试数据，第一行为 n（1<=n<=500000），表示 ip 地址库的 ip 区
间数量。
接下来的 n 行，每行包含三个数据，ip_start,   ip_end，province_id，表示 ip
区间的开始区间，结束区间，以及对应的省份 id。ip 区间为闭区间，而且各个
区间不会重叠交叉。ip 地址以 a.b.c.d 形式ᨀ供。
接下来是 m(1<=m<=500000)，表示用户访问日志数。
接下来的 m 行，每行是一个 ip 地址，表示用户访问的 ip 记录。
输出：
针对每组测试数据，按省份 id 输出每个省的访问数，每行第一个为省份 id，
第二个为该省份的访问数量，中间用空格分开。如果某个省份没有访问过，则访
问次数为 0。每组数据按省份 id 从小到大输出。
样例输入：
2
1
192.168.12.2 192.168.12.3 1
2
10.100.108.34
192.168.12.2
3
10.1.1.1 10.2.255.255 1
223.199.12.2 224.200.19.12 2
172.2.12.0 172.16.2.2 3
3
10.2.123.123
10.2.255.255
223.199.13.255
样例输出：
1   1
1   2
2   1
3   0
*/
