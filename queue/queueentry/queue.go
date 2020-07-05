package main

// A FIFO queue
type Queue2 []int

//Pushes the element into the queue
// eg. q.Push(123)
func (q *Queue2) Push(v int) {
	*q = append(*q, v)
}

//Pops element from head
func (q *Queue2) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

//Returns if the queue is empyt or not
func (q *Queue2) IsEmpty() bool {
	return len(*q) == 0
}
