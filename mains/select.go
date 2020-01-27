package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(
				time.Duration(rand.Intn(1500)) *
					time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d\n",
			id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)
	n := 0
	var activatevalue int
	var values []int
	tm := time.After(10 * time.Second)
	tk := time.Tick(time.Second)
	for {
		var activateworker chan<- int
		if len(values) > 0 {
			activateworker = worker
			activatevalue = values[0]
		}

		select {
		case n = <-c1:
			values = append(values, n)
		case n = <-c2:
			values = append(values, n)
		case activateworker <- activatevalue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("time out")
		case <-tk:
			fmt.Printf("the len : %d\n", len(values))
		case <-tm:
			fmt.Println("bye bye")
			return
		}
	}
}
