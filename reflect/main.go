package main

import (
	"fmt"
	"reflect"
)

type Calculate interface {
	Calc() int
}

type Addition struct {
	a int
	b int
}

func (this Addition) Calc() int {
	return this.a + this.b
}

type Subtraction struct {
	a int
	b int
}

func (this Subtraction) Calc() int {
	return this.a - this.b
}

type Operation struct {
	Symbol string
	a      int
	b      int
}

func (this Operation) Do() {
	var cal = new(Calculate)
	v := reflect.ValueOf(cal).Elem()
	fmt.Println(v)
}

func main() {
	var cal = new(Calculate)
	v := reflect.ValueOf(cal).Elem()
	fmt.Println(v)
}
