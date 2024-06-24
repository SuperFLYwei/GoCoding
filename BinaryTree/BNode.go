package main

import "fmt"

type BNode struct {
	Data  interface{}
	Left  *BNode
	Right *BNode
}

func NewBNode() *BNode {
	return &BNode{
		Data:  nil,
		Left:  nil,
		Right: nil,
	}
}

// MidOrderTraverse 中序遍历
func MidOrderTraverse(tree *BNode) {
	if tree == nil {
		return
	}
	MidOrderTraverse(tree.Left)
	fmt.Println("Data:", tree.Data)
	MidOrderTraverse(tree.Right)
}
