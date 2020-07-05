package main

import (
	"fmt"
	"strconv"
	"testing"
)

func ExampleQueue2_Pop() {
	q := Queue2{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	// Output:
	// 1
	// 2
	// false
	// 3
	// true
}

func ExampleQueue_Pop() {
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

func BenchmarkSprintf(b *testing.B) {
	q := NewQueue()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		node1 := &Node{Id: i, Name: "N:" + strconv.Itoa(i)}
		q.Push(node1)
		pop, _ := q.Pop().(*Node)
		fmt.Sprintf("%d: push:%v--pop:%v", i, &node1, &pop)
	}
}
