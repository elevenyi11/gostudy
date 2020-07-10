package main

import (
	"fmt"
)

func main() {

	seq := []string{"a"}
	// 指定删除位置
	index := 0
	// 查看删除位置之前的元素和之后的元素
	fmt.Println(seq[:index], seq[index+1:])
	// 将删除点前后的元素连接起来
	seq = append(seq[:index], seq[index+1:]...)
	fmt.Println(seq)

}

func testQueue3(){
	q := NewQueue3()

	node1 := &Node3{Id: 1, Name: "1111"}
	fmt.Println(&node1)
	q.Push(node1)
	node2 := &Node3{Id: 2, Name: "2222"}
	fmt.Println(&node2)
	q.Push(node2)
	q.Remove(node1)
	q.Remove(node2)
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	q.Remove(node2)
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
}

type Node struct {
	Id   int
	Name string
}

func testQueue(){
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
