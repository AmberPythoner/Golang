package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func euler() {
	fmt.Printf("%f\n", cmplx.Pow(math.E, 1i*math.Pi)+1)
	fmt.Printf("%f\n", cmplx.Exp(1i*math.Pi)+1)
}

func trangle() {
	a, b, c := 4, 3, 0
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
	fmt.Println(math.Round(3.88))
}

func consts() {
	const a, b = 3, 4
	c := math.Sqrt(a*a + b*b)
	fmt.Println(c)
}

func main() {
	consts()
}
