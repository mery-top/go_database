package bplustree

import "fmt"

type BPlusTreeNode struct {
	t        int
	keys     []int
	children []*BPlusTreeNode
	isLeaf   bool
	next     *BPlusTreeNode
}

type BPlusTree struct {
	root *BPlusTreeNode
	t    int
}

func NewNode(t int, isLeaf bool) *BPlusTreeNode {
	return &BPlusTreeNode{
		t:        t,
		keys:     []int{},
		children: []*BPlusTreeNode{},
		isLeaf:   isLeaf,
		next:     nil,
	}
}

func NewBPlusTree(t int) *BPlusTree {
	root := NewNode(t, true)
	return &BPlusTree{root: root, t: t}
}

func (tree *BPlusTree) Insert(key int) {
	if tree.root == nil {
		tree.root = NewNode(tree.t, true)
	}
	root := tree.root
	if len(root.keys) == 2*tree.t-1 {
		newRoot := NewNode(tree.t, false)
		newRoot.children = append(newRoot.children, root)
		tree.splitChild(newRoot, 0)
		tree.root = newRoot
	}
	tree.insertNonFull(tree.root, key)
}

func (tree *BPlusTree) insertNonFull(node *BPlusTreeNode, key int) {
	if node.isLeaf {
		i := len(node.keys) - 1
		for i >= 0 && key < node.keys[i] {
			i--
		}
		node.keys = append(node.keys[:i+1], append([]int{key}, node.keys[i+1:]...)...)
	} else {
		i := len(node.keys) - 1
		for i >= 0 && key < node.keys[i] {
			i--
		}
		i++

		// Ensure the child node exists and split if necessary
		if len(node.children[i].keys) == 2*tree.t-1 {
			tree.splitChild(node, i)
			if key > node.keys[i] {
				i++
			}
		}
		tree.insertNonFull(node.children[i], key)
	}
}

func (tree *BPlusTree) splitChild(parent *BPlusTreeNode, index int) {
	t := tree.t
	child := parent.children[index]
	newChild := NewNode(t, child.isLeaf)

	mid := t - 1
	// Split the parent node keys and children
	parent.keys = append(parent.keys[:index], append([]int{child.keys[mid]}, parent.keys[index:]...)...)
	parent.children = append(parent.children[:index+1], append([]*BPlusTreeNode{newChild}, parent.children[index+1:]...)...)

	// Move the right half of the child keys to the new child node
	newChild.keys = append([]int{}, child.keys[mid+1:]...)
	child.keys = child.keys[:mid]

	if !child.isLeaf {
		// Move the right half of the child children to the new child node
		newChild.children = append([]*BPlusTreeNode{}, child.children[mid+1:]...)
		child.children = child.children[:mid+1]
	} else {
		// Link the new child to the next leaf node
		newChild.next = child.next
		child.next = newChild
	}
}

func (tree *BPlusTree) Traverse() {
	node := tree.root

	for !node.isLeaf {
		node = node.children[0]
	}

	for node != nil {
		fmt.Println(node.keys)
		node = node.next
	}
}

