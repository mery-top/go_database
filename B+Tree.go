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
	if tree.root == nil{
		tree.root = newNode(tree.t, true)
	}
	root:= tree.root
	if len(root.keys) == 2*tree.t-1{
		newRoot := newNode(tree.t, false)
		newRoot.children = append(newRoot.children, root)
		tree.splitChild(newRoot, 0)
		tree.root = newRoot
	}
	tree.insertNonFull(tree.root, key)
}


func (tree * BPlusTree) insertNonFull(node *BPlusTreeNode, key int){
	if node.isLeaf{
		i := len(node.keys) -1
		for i>=0 && key < node.keys[i]{
			i--
		}
		node.keys = append(node.keys[:i+1], append([]int{key}, node.keys[i+1:]...)...)
	} else{
		i:= len(node.keys) -1
		for i>=0 && key < node.keys[i]{
			i--
		}
		i++

		if i>= len(node.children){
			node.children = append(node.children, newNode(tree.t, true))
		}

		if len(node.children[i].keys) == 2*tree.t-1{
			tree.splitChild(node,i)
			if key > node.keys[i] {
				i++
			}
			tree.insertNonFull(node.children[i], key)
		}
	}


}

func(tree *BPlusTree) splitChild(parent *BPlusTreeNode, index int){
	t:= tree.t
	child := parent.children[index]
	newChild:= newNode(t, child.isLeaf)

	if len(parent.children) == 0{
		parent.children = make([] *BPlusTreeNode, 1)
	}

	mid:= t-1
	parent.keys = append(parent.keys[:index], append([] int {child.keys[mid]}, parent.keys[index:]...)...)
	parent.children = append(parent.children[:index +1], append([] *BPlusTreeNode{newChild}, parent.children[index+1:]...)...)

	newChild.keys = append([]int{}, child.keys[mid+1:]...)
	child.keys = child.keys[:mid]

	if !child.isLeaf{
		newChild.children = append([] *BPlusTreeNode{}, child.children[mid+1:]...)
		child.children = child.children[:mid+1]
	} else{
		newChild.next = child.next
		child.next = newChild
	}
}
func (tree *BPlusTree) Traverse(){
	node:= tree.root

	for !node.isLeaf{
		node = node.children[0]
	}

	for node!= nil{
		fmt.Println(node.keys)
		node = node.next
	}
}

func main(){
	t:=3
	tree:= newBPlusTree(t)

	keys := []int {10, 20, 5, 6, 12, 30, 7, 17}
	for _, key := range keys{
		tree.Insert(key)
	}
	fmt.Println("B+ Tree Traversal")
	tree.Traverse()
}



