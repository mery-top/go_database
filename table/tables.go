package main
// Tables with queries
import (
	"go-db/bplustree"
	"fmt"
)

type Row struct{
	values []interface{}
}

type Table struct{
	name string
	columns []string
	rows []*Row
	primaryKeyIndex *bplustree.BPlusTree
}


func (table *Table) Insert(values []interface{}, primaryKey int){
	row:= &Row{values: values}
	
}


func main(){
	node:= bplustree.NewNode(3, true)
	fmt.Println(node)
}