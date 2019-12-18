package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename = "abc.txt"
	if contents, error := ioutil.ReadFile(filename); error != nil {
		fmt.Println(error)
	} else {
		fmt.Printf("%s\n", contents)
		fmt.Println(contents)
	}
}
