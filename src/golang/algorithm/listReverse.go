package main

import "fmt"

//import "container/list"

func main(){
	//l:=list.New()
	//l.PushBack(1)
	//l.PushBack(2)
	//l.PushBack(3)
	//l.PushBack(4)
	//l.PushBack(5)

	l1:=&ListNode{1,nil}
	l2:=&ListNode{2,l1}
	l3:=&ListNode{3,l2}
	l4:=&ListNode{4,l3}
	l5:=&ListNode{5,l4}
	res:=reverseList(l5)
	for true{
		if  res ==nil{
			break
		}
		fmt.Println(res.Val)
		res = res.Next
	}
}
/**
 * Definition for singly-linked list.
 *
 */

type ListNode struct {
	     Val int
	     Next *ListNode
	 }

func reverse(head *ListNode) *ListNode {

	var prev,tmp *ListNode
	for head!=nil{
		tmp = head.Next
		head.Next = prev
		prev=head
		if tmp==nil{
			break
		}
		head = tmp
	}
	return head
}

func  reverseList(head *ListNode)*ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}
	return prev
}
