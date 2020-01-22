package main

import (
	"fmt"
	"math"
)

func Trangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}
func main() {
	fmt.Println(Trangle(3, 4))
}
