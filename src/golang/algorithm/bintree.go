package algorithm

import (
)


func run1() {
	//tree
	r := NewTreeNode(3)
	r.L = NewTreeNode(2)
	r.R = NewTreeNode(5)

	//order
}

/*
type BinTree struct {
	root *TreeNode
	cur  *TreeNode
}
*/

type TreeNode struct {
	Data interface{}
	L    *TreeNode
	R    *TreeNode
	P    *TreeNode
}

func NewTreeNode(d interface{}) *TreeNode {
	return &TreeNode{Data: d}
}

func (*TreeNode) MidIterator() {

}
