package main

import (
	"fmt"
	"sync"
)

func doWork(id int, w worker) {
	for n := range w.in {
		fmt.Printf("the channer get the id %d, the values is %c \n", id, n)
		w.done()
	}
}

func createWork(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w)
	return w
}

type worker struct {
	in   chan int
	done func()
}

func channerDemo() {
	var wg sync.WaitGroup
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWork(i, &wg)
	}
	wg.Add(20)
	for i, w := range workers {
		w.in <- 'a' + i
	}

	for i, w := range workers {
		w.in <- 'A' + i
	}
	wg.Wait()
}

func main() {
	channerDemo()
	//bufferChanner()
}
