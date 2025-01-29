package main

import "fmt"

type BPlusTreeNode struct{
	t int
	keys []int
	children []*BPlusTreeNode
	isLeaf bool
	next *BPlusTreeNode
}

type BPlusTree struct{
	root *BPlusTreeNode
	t int
}

func newNode(t int, isLeaf bool) *BPlusTreeNode{
	return &BPlusTreeNode{
		t: t,
		keys: []int{},
		children: []*BPlusTreeNode{},
		isLeaf: isLeaf,
		next: nil,
	}
}

func newBPlusTree(t int) *BPlusTree{
	root:= newNode(t, true)
	return &BPlusTree{root: root, t: t}
}




