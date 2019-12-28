package main

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func CreateNode(value int) *Node {
	return &Node{
		Value: value,
		Left:  nil,
		Right: nil,
	}
}

func (root *Node) Print() {
	fmt.Println(root.Value)
}

func main() {
	root := CreateNode(10)
	root.Print()
}
