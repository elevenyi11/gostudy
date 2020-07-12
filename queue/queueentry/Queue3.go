package main

import (
	"sync"
)

type Node3 struct {
	Id   int32
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
	oldItem := q.last
	newItem := &leaguetopNode{prve: oldItem, item: item, next: nil}
	q.last = newItem
	q.n++
	if oldItem == nil {
		q.first = newItem
		return
	}
	newItem.prve = oldItem
	oldItem.next = newItem
}

// 参考 chan.go  func dequeue
func (q *Queue3) Pop() *Node3 {
	for {
		if q.first == nil {
			return nil
		}
		item := q.first
		nextItem := item.next
		if nextItem == nil {
			q.first = nil
			q.last = nil
		} else {
			nextItem.prve = nil
			q.first = nextItem
			item.next = nil // mark as removed
		}
		q.n--
		return item.item
	}
}

// 仅返回 不移除
func (q *Queue3) Peek() *Node3 {
	return q.first.item
}
func (q *Queue3) Remove(req *Node3) {
	if q.IsEmpty() {
		return
	}
	oldLast := &q.first
	for i := 0; i < q.n; i++ {
		currentQueue := *oldLast
		if i == 0 {
			currentQueue = *oldLast
		} else {
			currentQueue = currentQueue.next
		}
		if currentQueue.item.Id == req.Id {
			if i == 0 {
				q.first = currentQueue.next
				if q.first != nil {
					q.first.prve = nil
				}
			} else {
				(*oldLast).next = currentQueue.next
				(*currentQueue.next).prve = *oldLast
			}

			q.n--

			if i == q.n {
				q.last = *oldLast
				if q.last != nil {
					q.last.next = nil
				}
			}
			break
		}
		if i > 0 {
			oldLast = &currentQueue.next
		}
	}
	return
}
