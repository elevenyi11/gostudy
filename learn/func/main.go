package main

import "fmt"

type T struct {
	value int
}

func (m T) StayTheSame() {
	m.value = 3
}

func (m *T) Update() {
	m.value = 3
}

func main() {
	m := T{0}

	fmt.Println(m)

	m.StayTheSame()
	fmt.Println(m)
	m.Update() //值作为接收者（T） 不会修改结构体值，而指针 *T 可以修改
	fmt.Println(m)
}
