package main

import (
	"fmt"
)

type Currency int

const (
	USD Currency = iota //美元
	EUR                 //欧元
	GBP                 //英镑
	RMB                 //软妹
)

func main() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	q := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q)

	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB])

	compareArray()
}

func compareArray() {
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, b == c, a == c)
	//d := [3]int{1, 2}
	//fmt.Println(a == d)
}
