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
	fmt.Printf("%d  ", root.Value)
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

func (node *Tree) TraverseFunc(f func(tree *Tree)) {
	if node == nil {
		return
	}

	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

func (node *Tree) TraverseWithChannel() chan *Tree {
	out := make(chan *Tree)
	go func() {
		node.TraverseFunc(func(tree *Tree) {
			out <- tree
		})
		close(out)
	}()
	return out
}
