package main

import (
	"fmt"
	"strings"
)

func Hello() {
	defer func() {
		a := "hello world"
		fmt.Println(strings.Index(a, "oo"))
		r := recover()
		fmt.Println(111)
		if err, ok := r.(error); ok {
			fmt.Println("this is the errors %v", err)
		} else {
			panic(fmt.Sprintf("i do not know what to dos %v", err))
		}
	}()
	//panic(errors.New("this is the error"))
	//panic(123)
	b := 0
	fmt.Println(5 / b)
}

func main() {
	Hello()
}
