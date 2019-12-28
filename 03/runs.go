package main

import (
	"fmt"
)

func main() {
	a := "Yes 我爱慕课网!"
	fmt.Println(a)
	for i, ch := range a {
		fmt.Printf("( %d, %c)  ", i, ch)
	}
	fmt.Println("***********************")
	runes := []rune(a)
	fmt.Println(runes)
	for i, v := range []rune(a) {
		fmt.Printf("(%d , %c)  ", i, v)
	}
}
