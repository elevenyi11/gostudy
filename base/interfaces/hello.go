package interfaces

import "fmt"

type Human interface {
	sayHello()
}

type  Chinese struct {
	name string
}

type English struct {
	name string
}

func (c *Chinese) sayHello(){
	fmt.Println(c.name,"：你好啊")
}
func (e *English) sayHello(){
	fmt.Println(e.name,": Hello!!!")
}

