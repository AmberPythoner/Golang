package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type addrs func() int

func IAddr() addrs {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
func (g addrs) Read(
	p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("aa %d\n", next)

	// TODO: incorrect if p is too small!
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	var a addrs = IAddr()
	printFileContents(a)
}
