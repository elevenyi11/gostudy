package main

import "fmt"

type Student struct {
	Age           int
	Name, Address string
	Person
}

type Person struct {
	PhoneNumber int
}

func main() {
	stu := Student{
		Age:  18,
		Name: "xiao ming",
	}

	fmt.Println(stu)

	fmt.Println(Student{Age: 16, Name: "haha", Address: "chengdu", Person: Person{12321}})

	return
}
