package main

import "fmt"

func main() {
	var a []int
	b := []int{1, 2, 3, 4, 5, 6}
	a = append(a, b...)
	fmt.Println(a)
	fmt.Println(len(a), cap(a))

	fmt.Println("create slices")
	ac := []int{1, 2, 3, 4}
	fmt.Println(ac, len(ac), cap(ac))
	aa := make([]int, 10, 20)
	fmt.Println(aa, len(aa), cap(aa))

	fmt.Println("copy the slices")
	copy(aa, a)
	fmt.Println(aa)

	fmt.Printf("pop the slices\n")
	aa = append(aa[:4], aa[5:]...)
	fmt.Println(aa, len(aa), cap(aa))
}
