package main

import "fmt"

func arrayToTree(arr []int, start int, end int) *BNode {
	var root *BNode
	if end >= start {
		root = NewBNode()
		mid := (start + end + 1) / 2
		root.Data = arr[mid]
		root.Left = arrayToTree(arr, start, mid-1)
		root.Right = arrayToTree(arr, mid+1, end)
	}
	return root
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("数组", data)
	root := arrayToTree(data, 0, len(data)-1)
	MidOrderTraverse(root)
}
