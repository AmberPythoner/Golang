package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	a, b := 3, 6
	a, b = b, a
	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)

}
