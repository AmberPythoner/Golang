package main

import (
	"fmt"
	"testmodes/Nodes"
)

func main() {
	root := new(Nodes.Tree)
	fmt.Println(root)
	root2 := Nodes.Tree{Value: 100}
	fmt.Println(root2)
}
