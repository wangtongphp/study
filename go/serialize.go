/**
* Golang 实现PHP serialize()函数。
* @example
       encodeRes, err = phpserialize.Encode(mapData);
       decodeRes, err = phpserialize.Decode(string);
*/
package main

import (
	"fmt"

	"github.com/wulijun/go-php-serialize/phpserialize"
)

func main() {
	data3 := make(map[interface{}]interface{})
	data3["test"] = true
	data3[int64(0)] = int64(5)
	data3["flt32"] = float32(2.3)

	data2 := make(map[interface{}]interface{})
	data2["test"] = true
	data2[int64(0)] = int64(5)
	data2["flt32"] = float32(2.3)
	data2["int64"] = int64(45)
	data2["data3"] = data3

	data := make(map[interface{}]interface{})
	data["arr"] = data2
	data["3"] = "s\"tr'}e"
	data["g"] = nil

	var (
		result    string
		decodeRes interface{}
		err       error
	)

	if result, err = phpserialize.Encode(data); err != nil {
		fmt.Errorf("encode data fail %v, %v", err, data)
		return
	}
	fmt.Println(result) //a:3:{s:3:"arr";a:4:{s:4:"test";b:1;i:0;i:5;s:5:"flt32";d:2.299999952316284;s:5:"int64";i:45;}i:3;s:7:"s"tr'}e";s:1:"g";N;}

	php_s_str := `a:2:{s:3:"sub";a:2:{s:1:"c";i:3;i:99;s:1:"b";}s:2:"aa";i:88;}`
	if decodeRes, err = phpserialize.Decode(php_s_str); err != nil {
		fmt.Errorf("decode data fail %v, %v", err, result)
		return
	}
	fmt.Println(decodeRes) //map[sub:map[c:3 99:b] aa:88]

}
