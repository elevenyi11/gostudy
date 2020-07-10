package main

import (
	"fmt"
	"sync"
)

type Node3 struct {
	Id int32
	Name string
}

type Queue3 struct {
	first    *leaguetopNode
	last     *leaguetopNode
	n        int
	datalock sync.RWMutex
}

type leaguetopNode struct {
	prve *leaguetopNode
	item *Node3
	next *leaguetopNode
}

func NewQueue3() *Queue3 {
	return &Queue3{}
}

func (q *Queue3) IsEmpty() bool {
	return q.n == 0
}

func (q *Queue3) Size() int {
	return q.n
}

func (q *Queue3) Push(item *Node3) {
	oldlast := q.last
	q.last = &leaguetopNode{prve: oldlast}
	q.last.item = item
	q.last.next = nil
	if q.IsEmpty() {
		q.first = q.last
	} else {
		oldlast.next = q.last
	}
	q.n++
	if q.n == 2{
		q.last.prve = q.first
	}
	fmt.Println(item)
}

func (q *Queue3) Pop() *Node3 {
	if q.IsEmpty() {
		return nil
	}
	q.datalock.Lock()
	defer q.datalock.Unlock()
	item := q.first.item
	q.first = q.first.next
	q.first.prve = nil
	if q.IsEmpty() {
		q.last = nil
	}
	q.n--
	return item
}
func (q *Queue3) Peek() *Node3{
	return q.first.item
}
func (q *Queue3) Remove(req *Node3) {
	if q.IsEmpty() {
		return
	}
	q.datalock.Lock()
	defer q.datalock.Unlock()

	oldLast := *q.first
	fmt.Println(oldLast)
	currentQueue := *q.first
	for i:= 0; i < q.n; i++ {
		tempItem := currentQueue.item
		if tempItem.Id == req.Id{
			if currentQueue.next != nil{
				currentQueue.next.prve = &oldLast
			}
			oldLast.next = currentQueue.next
			q.n--
			if i == 0{
				q.first = currentQueue.next
				if q.first != nil{
					q.first.prve = nil
				}
			}
			if i == q.n{
				q.last = currentQueue.prve
				if q.last != nil{
					q.last.next = nil
				}
			}
			break
		}else{
			oldLast = currentQueue
		}

		currentQueue = *currentQueue.next
		if currentQueue.next == nil{
			break
		}
	}
	return
}