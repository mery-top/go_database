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


func (tree *BPlusTree) Insert(key int){
	root:= tree.root
	if len(root.keys) == 2*tree.t-1{
		newRoot := newNode(tree.t, false)
		newRoot.children = append(newRoot.children, root)
		tree.splitChild(newRoot, 0)
		tree.root = newRoot.next
	}
	tree.insertNonFull(tree.root, key)
}


func (tree * BPlusTree) insertNonFull(node *BPlusTreeNode, key int){
	if node.isLeaf{
		i := len(node.keys) -1
		for i>=0 && key < node.key[i]{
			i--
		}
		node.keys = append(node.keys[:i+1], append([]int{key}, node.keys[i+1:]...)...)
	} else{
		i:= len(node.keys) -1
		for i>=0 && key < node.keys[i]{
			i--
		}
		i++


		if len(node.children[i].keys) == 2*tree.t-1{
			tree.splitChild(node,i)
			if key > node.keys[i] {
				i++
			}
			tree.insertNonFull(node.children[i], key)
		}
	}


}



