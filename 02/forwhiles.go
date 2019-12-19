package main

import "fmt"

func sumArgs(values ...int) (sum int) {
	for i := range values {
		sum += values[i]
	}
	return sum
}

func ops(a, b int) {

}

func main() {
	fmt.Println(sumArgs(1, 2, 3, 4, 5))
}
