package main

import "fmt"

func changearrValue(slice []int) {
	slice[0] = 100
}

func main() {
	arrs := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arrs)
	changearrValue(arrs[3:])
	fmt.Println(arrs)
	fmt.Printf("%d, %d", len(arrs[4:]), cap(arrs[4:]))
}
