package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 100000; i++ {
		go func(i int) {
			for {
				fmt.Printf("hello go countine from %d\n", i)
			}
		}(i)
	}
	time.Sleep(time.Minute)
}
