package main

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
)

func main() {
	c := New()
	c.Add("192.168.1.1", "192.168.1.2", "192.168.1.3", "192.168.1.4")
	fmt.Println(c.Get("h|user|28882392"), c.Get("h|user|28882393"), c.Get("h|user|28882394"), c.Get("h|user|28882399"))
}

type CHash struct {
	replicas int
	keys     []int //sort hash slice
	hashMap  map[int]string
}

func New() *CHash {
	c := &CHash{
		replicas: 50,
	}
	c.hashMap = make(map[int]string)
	return c
}

func (c *CHash) Add(keys ...string) {
	for _, v := range keys {
		for ii := 0; ii < c.replicas; ii++ {
			h := int(crc32.ChecksumIEEE([]byte(strconv.Itoa(ii) + v)))
			c.keys = append(c.keys, h)
			c.hashMap[h] = v
		}
	}
	sort.Ints(c.keys)
	return
}

func (c *CHash) Get(key string) string {

	h := int(crc32.ChecksumIEEE([]byte(key)))
	idx := sort.Search(len(c.keys), func(i int) bool {
		return c.keys[i] >= h
	})
	if idx == len(c.keys) {
		idx = 0
	}
	return c.hashMap[c.keys[idx]]
}
