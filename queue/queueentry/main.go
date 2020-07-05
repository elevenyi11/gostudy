package main

import (
	"fmt"
)

func main() {
	q := NewQueue()

	node1 := &Node{Id: 1, Name: "1111"}
	fmt.Println(&node1)
	q.Push(node1)
	node2 := &Node{Id: 2, Name: "2222"}
	fmt.Println(&node2)
	q.Push(node2)
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())

}

type Node struct {
	Id   int
	Name string
}
