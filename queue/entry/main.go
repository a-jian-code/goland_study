package main

import (
	"fmt"
	"goland/queue"
)

func main() {
	q := queue.Queue{2}
	q.Push(4)
	q.Push(5)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
}
