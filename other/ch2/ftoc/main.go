package main

import (
	"fmt"
)

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fTocf(freezingF))
	fmt.Printf("%g°F = %g°C\n", boilingF, fTocf(boilingF))
	v := 1
	incr(&v)
	fmt.Println(incr(&v))
}
func fTocf(f float64) float64 {
	return (f - 32) * 5 / 9
}
func incr(p *int) int {
	*p++
	return *p
}
