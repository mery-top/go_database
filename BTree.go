package main

import (
	"fmt"
)

type node struct{
	t int
	keys []int
	children []int
	is_leaf bool
	next *node
}

func node_init(t int, keys []int,children []int, is_leaf bool ) node{
	return node{
		t:t,
		keys: keys,
		children: children,
		is_leaf: is_leaf,
		next: nil,
	}
}

func init(t int){
	var root int = node(t)
	return root
}

func insert(node, key int){
	
}
