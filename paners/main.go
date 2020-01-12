package main

import (
	"bufio"
	"fmt"
	"os"
)

func fibs() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func WriteFile(filename string) {
	f := fibs()
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		if PathError, ok := err.(*os.PathError); ok == true {
			fmt.Println(PathError.Op)
			fmt.Println(PathError.Path)
			fmt.Println(PathError.Err)
		} else {
			fmt.Println(err.Error())
			panic("no occth error")
		}
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	WriteFile("fib.txt")
}
