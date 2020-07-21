package queue

import (
	"container/list"
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

func TestList(t *testing.T){
	l := list.New()
	e4 := l.PushBack(4)
	e1:= l.PushFront(1)
	l.InsertBefore(3,e4)
	l.InsertAfter(2,e1)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func TestExampleQueue3(t *testing.T) {
	q := NewQueue3()

	node1 := &Node3{Id: 1, Name: "1111"}
	fmt.Println(&node1)
	q.Push(node1)
	node2 := &Node3{Id: 2, Name: "2222"}
	fmt.Println(&node2)
	q.Push(node2)

	node3 := &Node3{Id: 3, Name: "3333"}
	fmt.Println(&node3)
	q.Push(node3)

	fmt.Println(q.IsEmpty())
	q.Remove(node2)
	fmt.Println(q)
	q.Remove(node1)
	fmt.Println(q)
	q.Remove(node3)
	fmt.Println(q)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
}

func TestExampleQueue3_2(t *testing.T) {
	q := NewQueue3()

	node1 := &Node3{Id: 1, Name: "1111"}
	fmt.Println(&node1)
	q.Push(node1)
	node2 := &Node3{Id: 2, Name: "2222"}
	fmt.Println(&node2)
	q.Push(node2)

	node3 := &Node3{Id: 3, Name: "3333"}
	fmt.Println(&node3)
	q.Push(node3)

	fmt.Println(q.IsEmpty())
	q.Remove(node2)
	fmt.Println(q)
	fmt.Println(q.Pop())
	q.Remove(node1)
	fmt.Println(q)
	q.Remove(node3)
	fmt.Println(q)

	fmt.Println(q.Pop())
}

func BenchmarkSprintf(b *testing.B) {
	q := NewQueue()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		node1 := &Node3{Id: int32(i), Name: "N:" + strconv.Itoa(i)}
		q.Push(node1)
		pop, _ := q.Pop().(*Node3)
		fmt.Sprintf("%d: push:%v--pop:%v", i, &node1, &pop)
	}
}
