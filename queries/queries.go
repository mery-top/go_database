package main

import (
	"go-db/bplustree"
	"fmt"
)

func main(){
	node:= bplustree.NewNode(3, true)
	fmt.Println(node)
}