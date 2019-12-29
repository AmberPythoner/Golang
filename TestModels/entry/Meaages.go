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

//  表面上 看起来 是重载，其实只是魔法函数，并不是是重载
func (root *MyTree) Travels() {
	fmt.Println("This is MyTree Travels")
}

func main() {
	//内嵌结构赋值
	root := MyTree{&Nodes.Tree{Value: 200}}
	root.Left = Nodes.CreateTree(50)
	root.Right = Nodes.CreateTree(300)
	root.Travels()
	root.Tree.Travels()
	fmt.Println()
	root.PostOrder()

	//说明 不是重载
	//var BaseTree *root.Tree
	//BaseTree = root
}
