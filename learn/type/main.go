package main

import "fmt"

type City string

type Age int

type Width int
type Height int

func main() {

	city := City("cd")
	fmt.Println(city)
	middle := Age(13)
	if middle > 12 {
		fmt.Println("Middle is bigger than 12")
	}
	printAge(int(middle))

	w := Width(15)
	h := Height(50)
	result := int(w) / int(h)
	fmt.Println(result)
}

func printAge(age int) {
	fmt.Println("Age is ", age)
}
