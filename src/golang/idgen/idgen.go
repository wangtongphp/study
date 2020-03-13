package idgen

import (
	"encoding/binary"
	"errors"
	"math/rand"
	"net"
	"runtime"
	"strconv"
	"strings"
	"time"
)

/*
func main() {

	go genUID()

	select {}
}
*/

var ipLong uint32

func init() {

	// 由于IP运行期间不会变，故先缓存起来
	ip, err := getLocalIp()
	if err != nil {
		return
	}
	//fmt.Println("IP:", ip)

	ipLong, err = ip2long(ip)
	if err != nil {
		return
	}
	//fmt.Println("ipLong:", ipLong)
}

// base_convert(intval(microtime(true)) * 100000 + mt_rand(0, 99999), 10, 36)
// . base_convert(mt_rand(10000, 19999) * 100000000000000 + ip2long(gethostbyname(gethostname())) * 10000 + getmypid(), 10, 36);
func genUID() (res string, err error) {

	// 第一部分
	microTime := time.Now().UnixNano() / 1e3
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randNum1 := r.Intn(99)
	part1 := strconv.FormatInt(microTime*100+int64(randNum1), 36)

	// 第二部分
	randNum2 := r.Intn(9999) + 10000
	pid, err := getGoID()
	if err != nil {
		return
	}
	part2 := strconv.FormatInt(int64(randNum2)*100000000000000+int64(ipLong)*10000+int64(pid), 36)

	// 结果
	res = part1 + part2

	//fmt.Println(microTime, randNum1, part1, randNum2, ipLong, pid, part2, res)
	return res, nil
}

func ip2long(ipAddr string) (uint32, error) {

	ip := net.ParseIP(ipAddr)
	if ip == nil {
		return 0, errors.New("wrong ipAddr format")
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip), nil
}

func getLocalIp() (string, error) {

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", errors.New("ip get fail")
}

func getGoID() (int, error) {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		return id, err
	}
	return id, nil
}
