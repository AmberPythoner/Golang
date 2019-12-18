package main

import "fmt"

func consts() {
	const (
		aa = iota
		bb
		cc
		dd
	)
	fmt.Println(aa, bb, cc, dd)
	const (
		m = 1 << (10 * iota)
		mb
		kb
		tb
		pb
	)
	fmt.Println(m, mb, kb, tb, pb)
}

func main() {
	consts()
}
