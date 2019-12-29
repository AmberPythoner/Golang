package main

import (
	"fmt"
	"testmodes/Nodes"
)

type MyTree struct {
	*Nodes.Tree //不声明变量名 为内嵌结构
}

func (root *MyTree) PostOrder() {
	if root == nil || root.Tree == nil {
		return
	}
	left := MyTree{root.Left}
	right := MyTree{root.Right}
	left.PostOrder()
	right.PostOrder()
	root.Print()
}

func main() {
	root := MyTree{&Nodes.Tree{Value: 200}}
	root.Left = Nodes.CreateTree(50)
	root.Right = Nodes.CreateTree(300)
	root.Travels()
	fmt.Println()
	root.PostOrder()

}
