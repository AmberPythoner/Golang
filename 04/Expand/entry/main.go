package main

import (
	"expand/queue"
	"fmt"
)

func main() {
	q := queue.Queue{1}
	q.Append(2)
	q.Append(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Isempty())
	fmt.Println(q.Pop())
	fmt.Println(q.Isempty())
	q.Append(12121)
	fmt.Println(q.Isempty())
	fmt.Println(q.Pop())
	fmt.Println(q.Isempty())
}
