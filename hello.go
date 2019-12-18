package main

import (
	"fmt"
	"runtime"
)

var (
	aa = 11
	bb = "22"
)

func varValues() {
	a, b, c := "aa", 22, "cc"
	fmt.Println(a, b, c)
}

func main() {
	fmt.Println("Hello")
	fmt.Println(runtime.GOARCH)
	fmt.Println(aa, bb)
	varValues()
}
