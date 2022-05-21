package main

import "fmt"

type Student struct {
	Age           int
	Name, Address string
	Person
}

func (s *Student) Hello() {
	fmt.Printf("Student say hello haha")
}

type Person struct {
	PhoneNumber int
}

func (p *Person) Hello() {
	fmt.Printf("hello haha")
}

func main() {
	stu := Student{
		Age:  18,
		Name: "xiao ming",
	}

	fmt.Println(stu)

	fmt.Println(Student{Age: 16, Name: "haha", Address: "chengdu", Person: Person{12321}})
	stu.Hello()
	return
}
