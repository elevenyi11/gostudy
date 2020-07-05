package main

import (
	"fmt"
)

func main() {
	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i])
	}
	typeError()
	floatFunc()
}

func typeError() {
	var apples int32 = 1
	var oranges int16 = 2
	var compote int = int(apples) + int(oranges)
	fmt.Println(compote)
}

func floatFunc() {
	f := 3.141
	i := int(f)
	fmt.Println(f, i)
	f = 1.99
	fmt.Println(int(f))
}
