package main

import (
	"fmt"
)

func main() {
	var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := arr[2:9:10]
	fmt.Printf("a: %v,%v,%v\n", s, len(s), cap(s))
	a := [2]int{}
	fmt.Printf("a: %p\n", &a)

	fmt.Println(a)

	silce1 := make([]int, 0)
	silce2 := []int{}
	var silce3 []int
	var silce4 []int

	fmt.Printf("silce1:%v, %p\n", silce1 == nil, &silce1)
	fmt.Printf("silce2:%v, %p\n", silce2 == nil, &silce2)
	silce3 = append(silce3, 1, 2, 3)
	fmt.Printf("silce3:%v, %p\n", silce3 == nil, &silce3)
	fmt.Printf("silce4:%v, %p\n", silce4 == nil, &silce4)

	var peo People = Student{}
	think := "bitch"
	fmt.Println(peo.Speak(think))

}

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}
