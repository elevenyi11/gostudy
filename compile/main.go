package main

import "fmt"

func main()  {
	a,b:= test(2,1)
	fmt.Println(a, b)
}

//​go tool compile -S test.go > test.s
func test(i, j int) (int, int) {
	a:=i+ j
	b:=i- j
	return a,b
}

