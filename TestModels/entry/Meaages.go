package main

import (
	"testmodes/Nodes"
)

func main() {
	root := Nodes.CreateTree(100)
	root.Print()
	root.SetValue(200)
	root.Print()
	root.Left = Nodes.CreateTree(50)
	root.Right = Nodes.CreateTree(300)
	root.Travels()
}
