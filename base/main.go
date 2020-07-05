package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 3
	ss = "kk"
	bb = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func triangle() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func enums() {
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)

	const (
		B float64 = 1 << (iota * 10)
		KB
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)
	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(B, KB, MB, GB, TB, PB, EB, ZB, YB)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func euler() {
	fmt.Printf("%.3f\n", cmplx.Exp(math.Pi)+1)

}

func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)

	var c int
	c = int(math.Sqrt((a*a + b*b)))
	fmt.Println(filename, c)
}

func main() {
	enums()
	triangle()
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	euler()
	consts()
}
