package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := NewLRUCache(3)
	l.Set("a", "aa")
	l.Set("b", "bb")
	l.Set("c", "cc")

	fmt.Println(l.Get("a"))

	lis := New()
	lis.PushFront("111")
	lis.PushFront("222")
	lis.PushFront("333")
	fmt.Println(lis)
}

type ele struct {
	k string
	v interface{}
}

type LRUCache struct {
	ll         *list.List
	m          map[string]*list.Element
	maxEntries int
}

func NewLRUCache(maxEntries int) *LRUCache {
	return &LRUCache{
		ll:         list.New(),
		m:          make(map[string]*list.Element),
		maxEntries: maxEntries,
	}
}

func (l *LRUCache) Get(k string) (val interface{}, ok bool) {
	if l.m == nil {
		return
	}
	v, ok := l.m[k]
	if !ok {
		return
	}

	l.ll.MoveToFront(v)
	return v.Value.(*ele).v, true
}

func (l *LRUCache) Set(k string, v interface{}) bool {

	if val, ok := l.m[k]; ok {
		l.ll.MoveToFront(val)
		val.Value.(*ele).v = v
		return true
	}
	ele := l.ll.PushFront(&ele{k: k, v: v})
	l.m[k] = ele
	if l.ll.Len() > l.maxEntries {
		l.RemoveOldest()
	}
	return true
}

func (l *LRUCache) RemoveOldest() {
	if l.m == nil {
		return
	}
	e := l.ll.Back()
	if e == nil {
		return
	}
	l.ll.Remove(e)
	delete(l.m, e.Value.(*ele).k)
	return

}

// ------------------- list --------------------

type Element struct {
	prev, next *Element
	list       *List
	Value      interface{}
}

type List struct {
	root   Element //how pointer
	length int
}

func New() *List {
	var l = &List{}
	l.length = 0
	l.root.next = &l.root
	l.root.prev = &l.root
	return l
}

func (l *List) Back() *Element {
	return l.root.prev

}

func (l *List) PushFront(v interface{}) {
	l.insertBeforeEle(v, l.root.next)
	return
}

func (l *List) MoveToFront(e *Element) {
	tmp := l.remove(e)
	l.insertBeforeEle(tmp.Value.(*ele).v, l.root.next)
}

func (l *List) Remove(e *Element) {
	l.remove(e)
	return
}

func (l *List) insertBeforeEle(v interface{}, ele *Element) {

	e := &Element{Value: v, list: l}
	tmp := ele
	e.prev = ele.prev
	ele.prev.next = e
	ele.prev = e
	e.next = tmp
	l.length++
	return
}

func (l *List) remove(e *Element) *Element {
	tmp := e

	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil
	e.prev = nil
	e.list = nil
	e.Value = nil

	e.list.length--

	return tmp
}
