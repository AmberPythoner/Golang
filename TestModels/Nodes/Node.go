package Nodes

import "fmt"

type Tree struct {
	Value       int
	Left, Right *Tree
}

func CreateTree(value int) *Tree {
	return &Tree{
		Value: value,
	}
}

func (root *Tree) Print() {
	fmt.Println(root.Value)
}

func (root *Tree) SetValue(value int) {
	root.Value = value
}

func (root *Tree) Travels() {
	if root == nil {
		return
	}
	root.Left.Travels()
	root.Print()
	root.Right.Travels()
}
