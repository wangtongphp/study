package main

import (
	"fmt"
	"migo/utils"
	"os"
	"strconv"
	"strings"
)

func main() {
	ip := "210.73.15.0/24"
	sl := strings.Split(ip, "/")
	bi := strconv.FormatInt(utils.IP2Int(sl[0]), 2)
	prex_cnt, _ := strconv.Atoi(sl[1])
	binary_prex := utils.Substr(bi, 0, prex_cnt)
	fmt.Println(sl[0], bi, binary_prex)
	i_bi := strconv.FormatInt(utils.IP2Int("210.73.15.1"), 2)
	i_binary_prex := utils.Substr(i_bi, 0, prex_cnt)
	fmt.Println(i_binary_prex)
	os.Exit(1)
}
