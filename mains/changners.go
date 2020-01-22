package main

import (
	"fmt"
	"time"
)

func createChannels(id int) chan int {
	c := make(chan int)
	go func() {
		for n := range c {
			fmt.Printf("the channer get the id %d, the values is %c \n", id, n)
		}
	}()
	return c
}

func chanherDemo() {
	var channers [10]chan int
	for i := 0; i < 10; i++ {
		channers[i] = createChannels(i)
	}
	for i := 0; i < 10; i++ {
		channers[i] <- i + 'a'
	}
	time.Sleep(time.Millisecond)
}

func bufferChanner() {
	c := make(chan int, 3)
	go func() {
		for n := range c {
			fmt.Printf("the buffer get the value %d \n", n)
		}
	}()
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	time.Sleep(time.Millisecond)
}
func main() {
	//chanherDemo()
	bufferChanner()
}
