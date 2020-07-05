package main

import (
	"fmt"
	"math"
)

type tmp interface {
	area() float32
}

type geometry interface {
	tmp
	perim() float32
}
type rect struct {
	len, wid float32
}

func (r rect) area() float32 {
	return r.len * r.wid
}

func (r rect) perim() float32 {
	return 2 * (r.len + r.wid)
}

type circle struct {
	radius float32
}

func (c circle) area() float32 {
	return c.radius * c.radius * math.Pi
}

func (c circle) perim() float32 {
	return 2 * math.Pi * c.radius
}

func show(name string, param interface{}) {
	switch param.(type) {
	case geometry:
		fmt.Printf("area of %v is %v \n", name, param.(geometry).area())
		fmt.Printf("perim of %v is %v \n", name, param.(geometry).perim())
	default:
		fmt.Println("wrong type")
	}
}

func main() {
	rec := rect{
		len: 2,
		wid: 3,
	}

	show("rect", rec)

	cir := circle{
		radius: 1,
	}
	show("circle", cir)

	show("test", "test param")
}
