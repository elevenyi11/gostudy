package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + 2)
	}
}

func main() {
	a := adder2(0)

	for i := 0; i < 10; i++ {
		var s int
		s, a = a(i)
		fmt.Printf("0+1+2+ ... + %d = %d\n", i, s)
	}

	b := adder()
	b(10)
	fmt.Println(b(10))
}
