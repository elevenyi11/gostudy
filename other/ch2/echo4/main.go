package main

import (
	"flag"
	"fmt"
	"strings"

	"GoStudy/ch2/tempconv"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
	fmt.Println(gcd(2, 3))
	fmt.Println(fib(5))
	fmt.Printf("Brrrrr!%v\n", tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
	display()

}

func display() {
	args := []float64{50, 80, 100}

	for _, arg := range args {
		f := tempconv.Fahrenheit(arg)
		c := tempconv.Celsius(arg)
		fmt.Printf("%s = %s, %s =%s \n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
